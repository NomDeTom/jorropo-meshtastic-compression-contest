package main

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/Jorropo/meshtastic-compression-contest/arithcode"
	"github.com/Jorropo/meshtastic-compression-contest/unishox2"
	"google.golang.org/protobuf/encoding/protowire"
)

const (
	TEXT_MESSAGE_APP = 1
	NODEINFO_APP     = 4
	COMPRESSED       = 7
	// highBitNoCompressPortnum marks packets that intentionally skipped compression.
	// Regular meshtastic portnums are low values, so this gives a deterministic alias.
	highBitNoCompressPortnum = uint64(1 << 31)
)

func noCompressAliasPortnum(portnum uint64) (uint64, bool) {
	if portnum&highBitNoCompressPortnum != 0 {
		return 0, false
	}
	return portnum | highBitNoCompressPortnum, true
}

type compressorNum uint8

const (
	TEXT_MESSAGE_APP_COMPRESSOR compressorNum = iota + 1
	NODEINFO_COMPRESSOR
)

func explodePacketForPortnumPayloadSubstitution(substitute func(portnum uint64, from, to uint32, payload []byte) (newPortnum uint64, newPayload []byte, changed bool)) compressor {
	return func(input compressionInput) []byte {
		packet := input.data
		var portnum uint64
		var payload []byte
		var startOfPortnum, endOfPortnum, startOfPayload, endOfPayload int
		var parseCursor int
		for parseCursor < len(packet) {
			num, typ, n := protowire.ConsumeTag(packet[parseCursor:])
			if n < 0 {
				panic("invalid packet")
			}
			begin := parseCursor
			parseCursor += n

			switch num {
			case 1, 2:
				switch num {
				case 1:
					// portnum
					if typ != protowire.VarintType {
						panic("invalid portnum field type")
					}
					v, n := protowire.ConsumeVarint(packet[parseCursor:])
					if n < 0 {
						panic("invalid portnum value")
					}
					if uint64(uint32(v)) != v {
						panic("portnum value out of range")
					}
					portnum = v
					parseCursor += n
					startOfPortnum = begin
					endOfPortnum = parseCursor
				case 2:
					// payload
					if typ != protowire.BytesType {
						panic("invalid payload field type")
					}
					v, n := protowire.ConsumeBytes(packet[parseCursor:])
					if n < 0 {
						panic("invalid payload value")
					}
					payload = v
					parseCursor += n
					startOfPayload = begin
					endOfPayload = parseCursor
				}
			default:
				// Unknown field, skip it
				n := protowire.ConsumeFieldValue(num, typ, packet[parseCursor:])
				if n < 0 {
					panic("invalid field value")
				}
				parseCursor += n
			}
		}

		if portnum == 0 || payload == nil {
			panic("missing required fields")
		}

		newPortnum, newPayload, changed := substitute(portnum, input.from, input.to, payload)
		if !changed {
			return packet
		}

		newPacket := make([]byte, 0, 256)
		startFirst, endFirst, startLast, endLast := startOfPortnum, endOfPortnum, startOfPayload, endOfPayload
		if startFirst > startLast {
			startFirst, endFirst, startLast, endLast = startLast, endLast, startFirst, endFirst
		}
		newPacket = append(newPacket, packet[:startFirst]...)
		newPacket = append(newPacket, packet[endFirst:startLast]...)
		newPacket = append(newPacket, packet[endLast:]...)

		newPacket = protowire.AppendTag(newPacket, 1, protowire.VarintType)
		newPacket = protowire.AppendVarint(newPacket, uint64(newPortnum))

		newPacket = protowire.AppendTag(newPacket, 2, protowire.BytesType)
		newPacket = protowire.AppendBytes(newPacket, newPayload)

		if len(newPacket) >= len(input.data) {
			return input.data
		}

		return newPacket
	}
}

func compressPerPortnumTuned(portnum uint64, from, to uint32, payload []byte) (uint64, []byte, bool) {
	var compressedPortnum compressorNum
	var compressedPayload []byte
	switch portnum {
	case TEXT_MESSAGE_APP:
		portnum = COMPRESSED
		compressedPortnum = TEXT_MESSAGE_APP_COMPRESSOR
		var output [256]byte
		n := unishox2.CompressAlphaOnly(payload, output[:], "", 0)
		if n < 0 {
			log.Fatalf("unishox2 compression failed")
		}
		compressedPayload = output[:n]
	case NODEINFO_APP:
		portnum = COMPRESSED
		compressedPortnum = NODEINFO_COMPRESSOR
		var ok bool
		compressedPayload, ok = nodeinfoCompression2(from, payload)
		if !ok {
			return 0, nil, false
		}
	default:
		return 0, nil, false
	}

	return COMPRESSED, append([]byte{byte(compressedPortnum)}, compressedPayload...), true
}

// compressPerPortnumArithmeticImplicitJorropo uses original per-portnum Jorropo CDFs
// (trained on full dataset, not sampled)
func compressPerPortnumArithmeticImplicitJorropo(portnum uint64, from, to uint32, payload []byte) (uint64, []byte, bool) {
	_ = from
	_ = to

	cdf := arithmeticGetCDFJorropo(portnum)
	compressed := arithcode.Encode(payload, cdf)
	if compressed != nil && len(compressed) < len(payload) {
		// Compression is implicit when portnum is unchanged.
		return portnum, compressed, true
	}

	aliasPortnum, ok := noCompressAliasPortnum(portnum)
	if !ok {
		return 0, nil, false
	}

	// No-compress is explicit via aliased portnum, payload remains unchanged.
	return aliasPortnum, payload, true
}

// compressPerPortnumArithmeticImplicitTom uses Tom's trained 25% per-portnum CDFs
// (trained on 25% sampled dataset with fallback to global)
func compressPerPortnumArithmeticImplicitTom(portnum uint64, from, to uint32, payload []byte) (uint64, []byte, bool) {
	_ = from
	_ = to

	// Tom's approach: use 25% trained per-portnum CDFs
	cdf := arithmeticGetCDFTom(portnum)
	compressed := arithcode.Encode(payload, cdf)
	if compressed != nil && len(compressed) < len(payload) {
		// Compression is implicit when portnum is unchanged.
		return portnum, compressed, true
	}

	aliasPortnum, ok := noCompressAliasPortnum(portnum)
	if !ok {
		return 0, nil, false
	}

	// No-compress is explicit via aliased portnum, payload remains unchanged.
	return aliasPortnum, payload, true
}

// Legacy alias for backwards compatibility
func compressPerPortnumArithmeticImplicit(portnum uint64, from, to uint32, payload []byte) (uint64, []byte, bool) {
	return compressPerPortnumArithmeticImplicitJorropo(portnum, from, to, payload)
}

func nodeinfoCompression2(from uint32, payload []byte) ([]byte, bool) {
	// parse packet
	var id []byte
	var macaddress [6]byte
	var shortname []byte
	var hammode, isUnmessageable bool
	var longname []byte
	var eccpubkey [32]byte
	var hwmodel uint64
	var devicerole uint64

	user := payload
	for len(user) > 0 {
		num, typ, n := protowire.ConsumeTag(user)
		if n < 0 {
			return nil, false
		}
		user = user[n:]

		switch num {
		case 1: // id
			if typ != protowire.BytesType {
				return nil, false
			}
			id, n = protowire.ConsumeBytes(user)
			if n < 0 {
				return nil, false
			}
			user = user[n:]
		case 2: // long name
			if typ != protowire.BytesType {
				return nil, false
			}
			longname, n = protowire.ConsumeBytes(user)
			if n < 0 {
				return nil, false
			}
			user = user[n:]
		case 3: // short name
			if typ != protowire.BytesType {
				return nil, false
			}
			shortname, n = protowire.ConsumeBytes(user)
			if n < 0 {
				return nil, false
			}
			user = user[n:]
		case 4: // macaddress
			if typ != protowire.BytesType {
				return nil, false
			}
			v, n := protowire.ConsumeBytes(user)
			if n < 0 || len(v) != len(macaddress) {
				return nil, false
			}
			copy(macaddress[:], v)
			user = user[n:]
		case 5: // hw model
			if typ != protowire.VarintType {
				return nil, false
			}
			hwmodel, n = protowire.ConsumeVarint(user)
			if n < 0 {
				return nil, false
			}
			user = user[n:]
		case 6: // is licensed
			if typ != protowire.VarintType {
				return nil, false
			}
			v, n := protowire.ConsumeVarint(user)
			if n < 0 || (v != 0 && v != 1) {
				return nil, false
			}
			user = user[n:]
			hammode = protowire.DecodeBool(v)
		case 7: // device role
			if typ != protowire.VarintType {
				return nil, false
			}
			devicerole, n = protowire.ConsumeVarint(user)
			if n < 0 {
				return nil, false
			}
			user = user[n:]
		case 8: // ecc public key
			if typ != protowire.BytesType {
				return nil, false
			}
			v, n := protowire.ConsumeBytes(user)
			if n < 0 || len(v) != len(eccpubkey) {
				return nil, false
			}
			copy(eccpubkey[:], v)
			user = user[n:]
		case 9: // is unmessageable
			if typ != protowire.VarintType {
				return nil, false
			}
			v, n := protowire.ConsumeVarint(user)
			if n < 0 || (v != 0 && v != 1) {
				return nil, false
			}
			user = user[n:]
			isUnmessageable = protowire.DecodeBool(v)
		default:
			return nil, false
		}
	}

	// [l: example]
	// means an optional field in the byte stream with length l and content example
	// <l: example>
	// means an optional field in the string stream with length l and content example

	// Format:
	// ckhsssmi xullllll [l: unishox payload] [1: id length ] <←: id payload > [2: higher macaddress ] [4: lower macaddress ] <s: short name > <l: long name > [32: ecc public key ] [¿: varuint hwmodel ] [¿: varuint device role ]
	// header bitfields:
	//   0:
	//     i - 0 = id default from source, 1 = id present // TODO: ever since 7c5e2bc95acf81a0997169e7a4243d2a0af963e7 nodes do not care about this field, just don't bother with it ?
	//     m - 0 = lower macaddress default from source, 1 = lower macaddress present
	//     sss - length of the short name (7 means default from source)
	//     h - 1 = licensed ham mode, 0 = normal mode
	//     k - 0 = no ecc public key, 1 = ecc public key present
	//     c - unishox2 compression
	//         when using unixhox2, all strings are stored in the unishox decompressed payload in the order they appear bellow.
	//         when not using unishox2 the strings are mixed in the byte stream as bellow.
	//   ... - same format as the uncompressed version, but without the strings mixed in:
	//   1:
	//     llllll - length of the long name (63 means default from source)
	//     u - 1 = is unmessageable, 0 = normal
	//     x - reserved for future use

	var header [2]byte
	binaryWithoutUnishox := make([]byte, 0, 256)
	binaryWithUnishox := make([]byte, 0, 256)
	stringsWithUnishox := make([]byte, 0, 256)

	{
		if hammode {
			header[0] |= 1 << 5
		}

		if isUnmessageable {
			header[1] |= 1 << 6
		}

		defaultID := fmt.Sprintf("!%08x", from)
		if string(id) != defaultID {
			header[0] |= 1 << 0
			if len(id) >= 256 {
				return nil, false
			}
			binaryWithoutUnishox = append(binaryWithoutUnishox, byte(len(id)))
			binaryWithUnishox = append(binaryWithUnishox, byte(len(id)))
			binaryWithoutUnishox = append(binaryWithoutUnishox, id...)
			stringsWithUnishox = append(stringsWithUnishox, id...)
		}

		binaryWithoutUnishox = append(binaryWithoutUnishox, macaddress[0:2]...)
		binaryWithUnishox = append(binaryWithUnishox, macaddress[0:2]...)
		if binary.BigEndian.Uint32(macaddress[2:]) != from {
			header[0] |= 1 << 1
			binaryWithoutUnishox = append(binaryWithoutUnishox, macaddress[:]...)
			binaryWithUnishox = append(binaryWithUnishox, macaddress[:]...)
		}

		defaultShortname := fmt.Sprintf("%04x", from&0xffff)
		if string(shortname) != defaultShortname {
			if len(shortname) >= 7 {
				return nil, false
			}
			header[0] |= byte(len(shortname)) << 2
			binaryWithoutUnishox = append(binaryWithoutUnishox, shortname...)
			stringsWithUnishox = append(stringsWithUnishox, shortname...)
		} else {
			header[0] |= 7 << 2
		}

		defaultLongName := fmt.Sprintf("Meshtastic %04x", from&0xffff)
		if string(longname) != defaultLongName {
			if len(longname) >= 63 {
				return nil, false
			}
			header[1] |= byte(len(longname))
			binaryWithoutUnishox = append(binaryWithoutUnishox, longname...)
			stringsWithUnishox = append(stringsWithUnishox, longname...)
		} else {
			header[1] |= 63
		}

		if eccpubkey != [32]byte{} {
			header[0] |= 1 << 6
			binaryWithoutUnishox = append(binaryWithoutUnishox, eccpubkey[:]...)
			binaryWithUnishox = append(binaryWithUnishox, eccpubkey[:]...)
		}

		binaryWithoutUnishox = binary.AppendUvarint(binaryWithoutUnishox, hwmodel)
		binaryWithUnishox = binary.AppendUvarint(binaryWithUnishox, hwmodel)

		binaryWithoutUnishox = binary.AppendUvarint(binaryWithoutUnishox, devicerole)
		binaryWithUnishox = binary.AppendUvarint(binaryWithUnishox, devicerole)
	}

	var output [256]byte
	n := unishox2.CompressAlphaOnly(stringsWithUnishox, output[:], "", 0)
	if n < 0 {
		log.Fatalf("unishox2 compression failed")
	}
	unishoxPayload := output[:n]

	compressed := make([]byte, 0, 256)
	if unishoxIsSmaller := len(unishoxPayload)+len(header)+len(binaryWithUnishox) < len(header)+len(binaryWithoutUnishox); len(unishoxPayload) <= 127 && unishoxIsSmaller {
		compressed = append(compressed, header[:]...)
		compressed = append(compressed, unishoxPayload...)
		compressed = append(compressed, binaryWithUnishox...)
	} else {
		compressed = append(compressed, header[:]...)
		compressed = append(compressed, binaryWithoutUnishox...)
	}

	return compressed, true
}
