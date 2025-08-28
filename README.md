# Tiny Logger

A fast, lightweight, zero-dependency logging solution for Go applications that prioritizes performance and
simplicity.

Compatible with Go version 1.18.x and above

## ✅ Key Features

- 🪶 **Minimal Footprint**: The entire library weights around `100kB` (Comments and README included)
- ⚖️ **Zero Dependencies**: Pure Go implementation with no external dependencies
- 🚀 **Blazing Fast**: Custom JSON and YAML marshaling optimized for logging
    - Up to 1.4x faster JSON marshaling than `encoding/json`
    - Up to 5x faster YAML marshaling than `gopkg.in/yaml.v3`
    - Benchmark: up to `~640ns` per log entry (including JSON/YAML serialization) even less using the `Default encoder`.
- 🎨 **Color Support**: Built-in ANSI color support for terminal output
- 🔀 **Thread-Safe**: Concurrent-safe logging with atomic operations
- ⏱️ **Time-Optimized**: Efficient date/time handling with minimal allocations

## 📊 Benchmark Results

Data retrieved by executing the `./test/benchmark_test.go` file on my personal computer.

| Encoder             | Configuration      | ns/op | B/op | allocs/op |
|---------------------|--------------------|-------|------|-----------|
| **Default Encoder** | All Properties OFF | 468.6 | 80   | 1         |
|                     | All Properties ON  | 580.9 | 104  | 2         |
| **JSON Encoder**    | All Properties OFF | 513.3 | 80   | 1         |
|                     | All Properties ON  | 595.7 | 104  | 2         |
| **YAML Encoder**    | All Properties OFF | 566.0 | 80   | 1         |
|                     | All Properties ON  | 657.5 | 104  | 2         |

## 🎯 Use Examples

````go
logger := logs.NewLogger()
logger.Warn("my warning test") // stdout: 'WARN: my warning test'
logger.Info("my", "into", "test", 2) // stdout: 'INFO: my info test 2'
logger.Debug("hey", "check this", "debug") // stdout: 'DEBUG: hey check this debug'
logger.Error("here is the error") // stderr: 'ERROR: here is the error'

logger := logs.NewLogger().
   SetLogLvl(ll.WarnLvlName).
   EnableColors(true).
   AddTime(true).
   AddDate(true)

logger.Warn("This is my Warn log", "Test arg") // stdout: WARN[03/11/2024 18:35:43]: This is my Warn log Test arg

logger.SetEncoder(shared.JsonEncoderType)
logger.Debug("This is my Debug log", "Test arg") // stdout: {"level":"DEBUG","date":"03/11/2024","time":"18:35:43","message":"This is my Debug log Test arg"}

logger.AddTime(false)
logger.Debug("This is my Debug log", "Test arg") // stdout: {"level":"DEBUG","date":"03/11/2024","message":"This is my Debug log Test second arg"}
````

## 🤝 Contributing

Contributions are welcome, Here's how you can help:

### Setting Up Development Environment

1. Fork the repository
2. Clone your fork:
3. Create a new branch:

   ```bash
   git checkout -b feat/your-feature-name
   ```

### Development Guidelines

1. **Code Style**
    - Follow standard Go formatting (`go fmt`)
    - Use meaningful variable names
    - Add comments for non-obvious code sections
    - Write tests for new functionality

2. **Testing**
    - Run tests: `make test`
    - Run benchmarks: `make test-benchmark`
    - Ensure test coverage remains high, it can be checked using `make test-coverage`

3. **Performance**
    - Avoid unnecessary allocations
    - Use benchmarks to verify performance impact
    - Profile code changes when necessary

### Submitting Changes

1. Write clear, concise commit messages
2. Push to your fork
3. Submit a Pull Request with:
    - Description of changes
    - Benchmark results (if applicable)

## 💡 Why Tiny Logger?

- **Lightweight**: No external dependencies mean faster builds and smaller binaries
- **Performance**: Custom JSON marshaling optimized specifically for logging
- **Simplicity**: Clean API design with a minimal learning curve
- **Reliability**: Thoroughly tested with high test coverage
- **Maintainability**: Small, focused codebase makes it easy to understand and modify

Need help? Feel free to open an issue

## 📝 License

MIT License—see [LICENSE](https://mit-license.org/) file for details
