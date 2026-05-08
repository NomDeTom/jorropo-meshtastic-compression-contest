# Training Functions and CDF Generation

This folder contains documentation and utilities for training compression algorithms and generating CDF (Cumulative Distribution Function) tables.

## Overview

The meshtastic-compression-contest uses trained CDF tables for arithmetic coding compression. These are generated from packet payload data and are critical for performance tuning.

## Training Dataset Generation

The main compression contest supports generating training datasets when `generateTrainingDataset = true` in `main.go`:

- Individual packets are saved to `train/packets/` directory
- Per-portnum payloads are saved to `train/{PORTNUM}/` directories
- This allows analysis of compression performance across different message types

## CDF Dictionary Implementations

Two CDF strategies are implemented:

### 1. Jorropo Strategy (Original Full CDF)
- Function: `arithmeticGetCDFJorropo()` in `arith_tables.go`
- Uses: Per-portnum CDFs trained on complete dataset
- Compression: `compressPerPortnumArithmeticImplicitJorropo()` in `portnum_tunned.go`
- Best for: Achieving highest compression ratios when full training data is available

### 2. Tom Strategy (25% Sampled CDF with Fallback)
- Function: `arithmeticGetCDFTom()` in `arith_tables.go`  
- Uses: Per-portnum CDFs trained on 25% sampled data, fallback to global CDF
- Compression: `compressPerPortnumArithmeticImplicitTom()` in `portnum_tunned.go`
- Best for: Balancing compression with robustness to unseen data patterns

## Supported Portnums

Training and compression functions support these portnum identifiers:

- **1**: TEXT_MESSAGE_APP
- **3**: POSITION_APP
- **4**: NODEINFO_APP
- **5**: ROUTING_APP
- **65**: STORE_FORWARD_APP
- **67**: TELEMETRY_APP
- **70**: TRACEROUTE_APP
- **71**: NEIGHBORINFO_APP

## Training Data Sources

CDF tables are generated from:
- `freq_stats.json`: Frequency statistics from 25% sampled packets
- `freq_stats_25pct_with_fallback.json`: Combined 25% sample with global fallback frequencies (in meshcompress/)
- Static arrays in `arith_tables.go`: Pre-computed CDF lookup tables

## Key Functions

### In main.go
- `generateDatasetBin(filename string)`: Generates binary dataset for testing, optionally creates training data
- `generatePortnumSummaryTable()`: Creates markdown table of compression results per portnum

### In portnum_tunned.go
- `compressPerPortnumArithmeticImplicitJorropo()`: Compression using original full CDFs
- `compressPerPortnumArithmeticImplicitTom()`: Compression using 25% sampled CDFs

### In arith_tables.go
- `arithmeticGetCDFJorropo(portnum uint64)`: Retrieves Jorropo CDF for portnum
- `arithmeticGetCDFTom(portnum uint64)`: Retrieves Tom CDF for portnum with fallback

## Workflow

1. **Enable Training**: Set `generateTrainingDataset = true` in main.go
2. **Run Contest**: Execute `go run .` to generate training data in `train/` directory
3. **Analyze**: Review per-portnum payload distributions in `train/{PORTNUM}/` directories
4. **Optimize CDFs**: Use frequency analysis to update CDF tables in `arith_tables.go`
5. **Disable Training**: Set `generateTrainingDataset = false` before benchmarking
6. **Benchmark**: Run final competition against compressed datasets

## Files Reference

- `main.go`: Contest orchestration, contains training dataset generation code (lines 812-880)
- `portnum_tunned.go`: Per-portnum compression implementations with separate CDF strategies
- `arith_tables.go`: CDF table definitions and getter functions
- `../meshcompress/freq_stats_25pct_with_fallback.json`: Sample frequency statistics
