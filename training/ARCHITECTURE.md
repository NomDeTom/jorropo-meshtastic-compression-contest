# Training Architecture Guide

## Overview

The training module provides utilities and documentation for:
1. **Generating training datasets** from packet payloads
2. **Understanding CDF strategies** for different compression implementations
3. **Managing training data workflows** during algorithm development

## File Structure

```
training/
├── README.md                    # Main documentation
├── ARCHITECTURE.md              # This file - design patterns and workflows
└── training_utils.go           # Reference implementation (in main package)
```

## Design Patterns

### 1. Portnum-Based CDF Training

Different message types (portnums) benefit from different CDF dictionaries:

```
Dataset → Frequency Analysis → Per-Portnum CDFs → Compression
         (per portnum)           (stored in       (improved
                                 arith_tables.go)  ratios)
                                       ↓
                                   Global Fallback
                                   (for robustness)
```

### 2. CDF Strategy Selection

Choose your CDF strategy based on use case:

| Strategy | Function | Best For | Tradeoff |
|----------|----------|----------|----------|
| **Jorropo** | `arithmeticGetCDFJorropo()` | Maximum compression | Requires full training data |
| **Tom** | `arithmeticGetCDFTom()` | Production robustness | Slightly larger output |

Both strategies support the same portnum compression functions:
- `compressPerPortnumArithmeticImplicitJorropo()`
- `compressPerPortnumArithmeticImplicitTom()`

### 3. Training Workflow

**Step 1: Enable Training**
```go
const generateTrainingDataset = true  // in main.go
```

**Step 2: Run Collection**
```bash
go run . 2>&1 | tee training_run.log
```

This generates:
- `train/packets/` - Individual packet binaries
- `train/1/`, `train/3/`, etc. - Per-portnum payloads
- `README.md` - Competition results

**Step 3: Analyze Results**
```bash
# Count training samples per portnum
for d in train/*/; do echo "$d: $(ls $d | wc -l) samples"; done
```

**Step 4: Generate Frequency Statistics**
Analyze payload data to compute CDF tables for each portnum.

**Step 5: Update CDF Tables**
Edit `arith_tables.go`:
- Update `arithmeticCDFByPortnum` with new CDFs
- Or extend `arithmeticCDFGlobal` for fallback strategy

**Step 6: Disable Training and Benchmark**
```go
const generateTrainingDataset = false  // in main.go
go build && ./meshtastic-compression-contest
```

## Integration with Contest Pipeline

The main contest (`main.go`) automatically handles training when enabled:

1. **Data Loading**: `generateDatasetBin()` reads SQLite packet data
2. **Training Data Collection**: When `generateTrainingDataset = true`:
   - Duplicates packets to `train/packets/` directory
   - Extracts payloads to portnum-specific directories
   - Maintains counters for statistics
3. **Temporary File Management**: Training data written to `/tmp` during runs
4. **Cleanup**: Training directories persist after run for analysis

## Key Decision Points

### When to Train

- ✅ DO train when adding new message types or compression algorithms
- ✅ DO train when dataset changes significantly
- ❌ DON'T train for final benchmarks (use pre-trained CDFs)

### Storage Locations

- **Training data**: Generated in `train/` during execution (can be removed after analysis)
- **CDF tables**: Embedded in `arith_tables.go` as Go arrays
- **Frequency data**: Reference copies in `../meshcompress/freq_stats_*.json`

## Common Tasks

### Add a New Portnum to Training

1. Ensure portnum is defined in `portnum.py` (meshcompress/)
2. Add mapping to `portnumFriendlyName()` in main.go
3. Add CDF entry to `arithmeticCDFByPortnum` in arith_tables.go
4. Run with `generateTrainingDataset = true`
5. Analyze results in `train/{NEW_PORTNUM}/` directory

### Compare Compression Strategies

The per-portnum table in `README.md` shows all strategies side-by-side:

```markdown
| Compressor | **Portnum 1**<br/>(TEXT_MESSAGE_APP) | **Portnum 3**<br/>(POSITION_APP) | ...
|---|---|---|
| arithmetic_Jorropo | 0.6234 | 0.8912 | ...
| arithmetic_Tom | 0.6245 | 0.8934 | ...
```

Lower ratio = better compression. Differences highlight CDF strategy effectiveness.

### Reproduce Training Results

To replicate a specific training run:

```bash
# Use specific dataset
sqlite3 meshtastic.db "SELECT * FROM packets" > packets.sql

# Enable training
vim main.go  # const generateTrainingDataset = true

# Run with timing
time go run . 2>&1 | tee training_$(date +%s).log

# Verify training data generated
ls -lah train/ | head -20

# Disable training for production
vim main.go  # const generateTrainingDataset = false
```

## Troubleshooting

### Training Files Not Generated

**Problem**: `train/` directory empty after run
**Solution**:
1. Verify `generateTrainingDataset = true` in main.go
2. Check disk space: `df -h`
3. Verify write permissions: `touch train/test && rm train/test`
4. Check database has packets: `sqlite3 meshtastic.db "SELECT COUNT(*) FROM packets;"`

### CDF Tables Not Updating

**Problem**: Modified `arithmeticCDFByPortnum` but no difference in results
**Solution**:
1. Verify function is called: Add debug print to `arithmeticGetCDFJorropo()`
2. Rebuild: `go clean && go build`
3. Check portnum values are correct (use `portnumFriendlyName()` to verify)
4. Verify CDF array has 257 elements (0-256 byte values)

### Memory Issues During Training

**Problem**: Out of memory with large dataset
**Solution**:
1. Use `/tmp` for training data (configured in `generateDatasetBin()`)
2. Split dataset into smaller chunks
3. Clean up old training directories: `rm -rf train/`
4. Monitor memory: `watch -n 1 'free -h'`

## Future Improvements

- [ ] Automatic CDF table generation from training data
- [ ] Statistical analysis of CDF effectiveness per portnum
- [ ] Incremental training (append to existing CDFs)
- [ ] Cross-validation of CDF strategies on holdout data
- [ ] Integration with CI/CD pipeline for automated training
