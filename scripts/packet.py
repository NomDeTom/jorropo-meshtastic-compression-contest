from dataclasses import dataclass
from portnum import PortNum

@dataclass
class Packet:
    header: bytes
    payload: bytes

    
    @property
    def id_to(self) -> int:
        return int.from_bytes(self.header[0:4], "big")
    
    @property
    def id_from(self) -> int:
        return int.from_bytes(self.header[4:8], "big")
    
    @property 
    def packet_id(self) -> int:
        return int.from_bytes(self.header[8:12], "big")
    
    @property
    def flags(self) -> int:
        return self.header[13]
    
    @property
    def hop_limit(self) -> int:
        return self.flags & 0x07
    
    @property
    def channel(self) -> int:
        return self.header[14]
    
    @property
    def next_hop(self) -> int:
        return self.header[15]

    @property
    def relayer_id(self) -> int:
        return self.header[16]
    
    @property
    def portnum(self) -> PortNum:
        value = self.parse_varint(self.payload)[0]
        if value > 256: return PortNum.PRIVATE_APP
        mapped = PortNum(value)
        return mapped
    
    @property
    def payload_str(self) -> str:
        return self.payload.decode(errors="ignore")

    def extract_data_payload(self) -> bytes | None:
        """Extract the inner payload bytes from the protobuf Data message (field 2)."""
        data = self.payload
        i = 0
        while i < len(data):
            tag, i = Packet._parse_varint(data, i)
            field_number = tag >> 3
            wire_type = tag & 0x07
            if wire_type == 0:  # varint - skip
                _, i = Packet._parse_varint(data, i)
            elif wire_type == 2:  # length-delimited
                length, i = Packet._parse_varint(data, i)
                if field_number == 2:
                    return data[i:i + length]
                i += length
            elif wire_type == 1:  # 64-bit fixed
                i += 8
            elif wire_type == 5:  # 32-bit fixed
                i += 4
            else:
                break
        return None


    @staticmethod
    def parse_varint(data: bytes, offset: int = 0):
        key = Packet._parse_varint(data, offset)
        value = Packet._parse_varint(data, offset+1)
        return value

    @staticmethod
    def _parse_varint(data: bytes, offset = 0) -> tuple[int, int]:
        result = 0
        shift = 0

        while True:
            parsing_byte = data[offset]
            offset += 1

            result |= (parsing_byte & 0x7f) << shift
            
            # Check for continuation bit
            if (parsing_byte & 0x80) == 0:
                break
            
            shift += 7
            if shift >= 64:
                raise Exception("Max varint bit length of 64 exceeded.")
        return result, offset
        
        
        """Modified from https://github.com/protocolbuffers/protobuf/blob/main/python/google/protobuf/internal/decoder.py"""
        result = 0
        shift = 0
        buffer = bytearray(pl_bytes)
        while True:
            if pos is None:
                # Read from BytesIO
                try:
                    b = buffer.pop()
                except IndexError as e:
                    if shift == 0:
                    # End of BytesIO.
                        return None
                    else:
                        raise ValueError('Fail to read varint %s' % str(e))
            else:
                b = buffer[pos]
                pos += 1
            result |= ((b & 0x7f) << shift)
            if not (b & 0x80):
                result = int(result)
                return result if pos is None else (result, pos)
            shift += 7
            if shift >= 64:
                raise Exception('Too many bytes when decoding varint.')


    @staticmethod
    def from_hex(in_str: str) -> "Packet":
        b = bytes.fromhex(in_str)
        h = b[:16]
        p = b[16:]
        return Packet(h, p)

if __name__ == "__main__":
    
    assert Packet.parse_varint(bytes.fromhex("08 05")) == (5,2)
    assert Packet.parse_varint(bytes.fromhex("08 85 01")) == (133,3)
    assert Packet.parse_varint(bytes.fromhex("08 43")) == (67, 2)
    
    test_hex = "08 78 eb a2  98 70 c3 08  ab 44 bd 2f  a0 08 00 00    08 05 12 02 18 08 35 be a9 aa fe 48 01 "
    test_hex = "10 af 6d fa  e0 db 3d 43  3c 67 d5 78  e0 08 00 00    08 05 12 02 18 08 35 3a 40 45 91 48 01"
    packet = Packet.from_hex(test_hex)
    ascii = bytes.fromhex(test_hex).decode(errors="ignore")
    print(test_hex)
    print(repr(ascii))
    print(f"     To: {packet.id_to}")
    print(f"   From: {packet.id_from}")
    print(f"pack_id: {packet.packet_id}")
    print(f"PortNum: {packet.portnum}")
    pass