# Data Directory

This directory contains packet data and helper scripts for the meshtastic compression contest.

## Structure

### `packets/` - Packet Data and Databases
- **packetlog.txt** - Original packet log (40MB, source data)
- **packetlog_small.txt** - Smaller test dataset (270KB, for quick iteration)
- **packets_recovered.db** - SQLite database of decoded packets (15MB, used by main.go)
- **packets_recovered_small.db** - Smaller test database (112KB)
- **packets.bin** - Binary format packets (can be regenerated)

### Data Flow
```
packetlog.txt ──convert_to_sqlite.py──> packets_recovered.db ──main.go──> compression results
                                                                 └──> README.md (final record)
```

## Helper Scripts (in `../scripts/`)

- **convert_to_sqlite.py** - Converts packetlog.txt to SQLite database
- **convert_packetlog_to_binary.py** - Converts packetlog.txt to binary format
- **packet.py** - Packet parsing utility
- **portnum.py** - Port number definitions

## Usage

### Regenerating the SQLite Database
```bash
cd ../scripts
python3 convert_to_sqlite.py
```

### Running the Contest
```bash
cd /home/foxbox/meshtastic-compression-contest
go run main.go portnum_tunned.go utils.go
```

The contest will:
1. Load packets from `data/packets/packets_recovered.db`
2. Generate temporary dataset files in `/tmp/`
3. Write final results to `README.md`
4. Clean up temporary files automatically

## Important Notes

- **README.md is the final record** of benchmark results - preserve it
- **Temporary files** are written to `/tmp/` during execution, not in the repo
- **Database paths** in main.go point to `data/packets/packets_recovered.db`
- Old result files (.txt, .csv, .json) are automatically cleaned up before each run
