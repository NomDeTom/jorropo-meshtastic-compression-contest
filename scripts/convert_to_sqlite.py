#!/usr/bin/env python3
"""Convert packetlog.txt to SQLite database for Go compression contest."""

import sqlite3
import struct
from pathlib import Path
from packet import Packet

SOURCE_FILE = Path("packetlog.txt")
DB_PATH = Path("packets_recovered.db")

def encode_varint(value: int) -> bytes:
    """Encode a protobuf varint."""
    result = []
    while value > 0x7f:
        result.append((value & 0x7f) | 0x80)
        value >>= 7
    result.append(value & 0x7f)
    return bytes(result)

def encode_fixed32(value: int) -> bytes:
    """Encode a protobuf fixed32 (little-endian)."""
    return struct.pack('<I', value)

def create_wrapper_message(from_id: int, to_id: int, data: bytes) -> bytes:
    """Create a protobuf message with from/to/data fields.
    
    Field 1: from (fixed32) - tag 0x0d (1 << 3 | 5)
    Field 2: to (fixed32) - tag 0x15 (2 << 3 | 5)
    Field 4: data (bytes) - tag 0x22 (4 << 3 | 2)
    """
    msg = b''
    
    # Field 1: from (fixed32)
    msg += b'\x0d'  # tag
    msg += encode_fixed32(from_id)
    
    # Field 2: to (fixed32)
    msg += b'\x15'  # tag
    msg += encode_fixed32(to_id)
    
    # Field 4: data (bytes, length-delimited)
    msg += b'\x22'  # tag
    msg += encode_varint(len(data))
    msg += data
    
    return msg

def main():
    print(f"Reading packets from {SOURCE_FILE}")
    lines = SOURCE_FILE.read_text().splitlines()
    packets = [Packet.from_hex(line) for line in lines]
    print(f"Read {len(packets)} packets")
    
    # Remove old database if it exists
    if DB_PATH.exists():
        DB_PATH.unlink()
    
    # Create database
    conn = sqlite3.connect(str(DB_PATH))
    cursor = conn.cursor()
    
    # Create packet table
    cursor.execute('''
        CREATE TABLE packet (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            payload BLOB NOT NULL
        )
    ''')
    
    # Insert packets
    valid_count = 0
    for i, packet in enumerate(packets):
        if i % 100 == 0:
            print(f"Processing {i}/{len(packets)}...")
        
        # The payload contains the Data message (portnum + inner payload)
        # We wrap it with from/to/data to create the expected structure
        if len(packet.payload) > 0 and len(packet.payload) <= 255:
            # Create wrapper message with from/to/data
            wrapped = create_wrapper_message(packet.id_from, packet.id_to, packet.payload)
            cursor.execute('INSERT INTO packet (payload) VALUES (?)', (wrapped,))
            valid_count += 1
    
    conn.commit()
    
    # Verify
    cursor.execute('SELECT COUNT(*) FROM packet')
    count = cursor.fetchone()[0]
    print(f"Inserted {count} valid packets into database")
    
    conn.close()
    print(f"Database created at {DB_PATH}")

if __name__ == "__main__":
    main()
