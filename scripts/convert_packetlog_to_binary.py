#!/usr/bin/env python3
"""Convert packetlog.txt to packets.bin format for Go compression contest."""

import struct
from pathlib import Path
from packet import Packet

PACKET_SOURCE = Path("packetlog.txt")
BINARY_DEST = Path("packets.bin")

def parse_varint(data: bytes) -> tuple[int, int]:
    """Parse a protobuf varint and return (value, bytes_read)."""
    result = 0
    shift = 0
    for i, byte in enumerate(data):
        result |= (byte & 0x7f) << shift
        if (byte & 0x80) == 0:
            return result, i + 1
        shift += 7
    return 0, 0

def packets_from_file(path: Path) -> list[Packet]:
    print(f"Reading packet data from {path}")
    source_lines = path.read_text().splitlines()
    parsed = [Packet.from_hex(d) for d in source_lines]
    print(f"Read {len(parsed):,} packets from file.")
    return parsed

def extract_lora_payload(packet_payload: bytes) -> tuple[int, int, bytes, bool]:
    """Extract from, to, and data from a LoRa Data message payload.
    
    This mimics the Go extractLoraPayloadFromMessage function.
    """
    pos = 0
    from_id = None
    to_id = None
    data = None
    
    while pos < len(packet_payload):
        # Read tag (field number and wire type)
        tag, bytes_read = parse_varint(packet_payload[pos:])
        pos += bytes_read
        
        field_num = tag >> 3
        wire_type = tag & 0x7
        
        if field_num == 1:  # from field (fixed32)
            if wire_type != 5:
                return 0, 0, b'', False
            from_id = struct.unpack('<I', packet_payload[pos:pos+4])[0]
            pos += 4
        elif field_num == 2:  # to field (fixed32)
            if wire_type != 5:
                return 0, 0, b'', False
            to_id = struct.unpack('<I', packet_payload[pos:pos+4])[0]
            pos += 4
        elif field_num == 4:  # data field (bytes)
            if wire_type != 2:
                return 0, 0, b'', False
            length, bytes_read = parse_varint(packet_payload[pos:])
            pos += bytes_read
            data = packet_payload[pos:pos+length]
            pos += length
        else:
            # Skip unknown fields
            if wire_type == 0:  # varint
                _, bytes_read = parse_varint(packet_payload[pos:])
                pos += bytes_read
            elif wire_type == 2:  # bytes
                length, bytes_read = parse_varint(packet_payload[pos:])
                pos += bytes_read + length
            elif wire_type == 5:  # fixed32
                pos += 4
            elif wire_type == 1:  # fixed64
                pos += 8
            else:
                return 0, 0, b'', False
    
    if from_id is not None and to_id is not None and data is not None:
        return from_id, to_id, data, True
    return 0, 0, b'', False

def convert_to_binary(output_path: Path):
    """Convert packetlog.txt to packets.bin format."""
    packets = packets_from_file(PACKET_SOURCE)
    
    valid_packets = []
    for i, packet in enumerate(packets):
        from_id, to_id, data, ok = extract_lora_payload(packet.payload)
        if not ok or len(data) == 0 or len(data) > 255:
            continue
        valid_packets.append((from_id, to_id, data))
    
    print(f"Using {len(valid_packets)} valid packets (skipped {len(packets) - len(valid_packets)})")
    
    with open(output_path, 'wb') as f:
        # Write total number of rows (uint64, little-endian)
        num_rows = len(valid_packets)
        f.write(struct.pack('<Q', num_rows))
        
        # Write each packet
        for i, (from_id, to_id, data) in enumerate(valid_packets):
            if i % 5000 == 0:
                print(f"Writing packet {i}/{num_rows}...")
            
            # Write from (uint32, little-endian)
            f.write(struct.pack('<I', from_id))
            
            # Write to (uint32, little-endian)
            f.write(struct.pack('<I', to_id))
            
            # Write payload length (uint8)
            f.write(struct.pack('B', len(data)))
            
            # Write payload
            f.write(data)
    
    # Get file size
    file_size = output_path.stat().st_size
    print(f"Wrote {file_size/1024/1024:,.1f}MB to {output_path}")
    return output_path

if __name__ == "__main__":
    convert_to_binary(BINARY_DEST)

