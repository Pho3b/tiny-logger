# Tiny Logger

A fast, lightweight, zero-dependency logging solution for Go applications that prioritizes performance and
simplicity.    
Compatible with Go version 1.18.x and above

"TinyLogger is the fastest way to log loosely typed data (interfaces)"

## ✅ Key Features

- **Lightweight**: No external dependencies mean faster builds and smaller binaries <br>
  The dependencies that you see in the `go.mod` file are not included in the final binary since they are only used in `_test` files.
- **Simplicity**: Clean API design with a minimal learning curve. You'll set it up in seconds.
- **Performance**: The library is benchmarked to be very fast. It implements custom JSON and YAML marshaling
  specifically optimized for logging
    - Up to 1.4x faster JSON marshaling than `encoding/json`
    - Up to 5x faster YAML marshaling than `gopkg.in/yaml.v3`
- **Color Support**: Built-in ANSI color support for terminal output
- **Thread-Safe**: Concurrent-safe logging with atomic operations
- **Time-Optimized**: Efficient date/time print built-int logic with minimal allocations
- **Reliability**: Thoroughly tested with high test coverage
- **Maintainability**: A small, focused codebase makes it easy to understand and modify at will

## 🎯 Use Examples

````go
/******************** Basic Logging methods usage ********************/
logger := logs.NewLogger()
logger.Warn("my warning test") // stdout: 'WARN: my warning test'
logger.Info("my", "info", "test", 2) // stdout: 'INFO: my info test 2'
logger.Debug("hey", "check this", "debug") // stdout: 'DEBUG: hey check this debug'
logger.Error("here is the error") // stderr: 'ERROR: here is the error'

/******************** Configuration setup example ********************/
logger := logs.NewLogger().
    SetLogLvl(ll.WarnLvlName).
    EnableColors(true).
    AddTime(true).
    AddDate(true)

logger.Warn("This is my Warn log") // stdout: WARN[03/11/2024 18:35:43]: This is my Warn log

logger.SetEncoder(shared.JsonEncoderType)
logger.Debug("This is my Debug log") // stdout: {"level":"DEBUG","datetime":"03/11/2024 18:35:43","message":"This is my Debug log"}

logger.AddTime(false)
logger.Debug("This is my Debug log") // stdout: {"level":"DEBUG","date":"03/11/2024","message":"This is my Debug log"}

/******************** Logging to a file example ********************/
file, err := os.OpenFile("./my-out-file.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
if err != nil {
    println("ERROR: cannot open out file", err)
}

logger := logs.NewLogger().SetLogFile(file) // From this point on loggers logs will be redirected to the file
logger.Debug("This is my Debug log", "Test arg")

logger.CloseLogFile()

/******************** Date Time Formatting example ********************/
logger := logs.NewLogger().SetDateTimeFormat(shared.IT)
logger.Debug("This is my Debug log", "Test arg") // stdout: 03/11/2024 18:35:43: This is my Debug log Test arg
logger.SetDateTimeFormat(shared.US)
logger.Debug("This is my Debug log", "Test arg") // stdout: 11/03/2024 18:35:43: This is my Debug log Test arg
logger.SetDateTimeFormat(shared.UnixTimestamp)
logger.Debug("This is my Debug log", "Test arg") // stdout: 1690982143.000000 This is my Debug log Test arg
````

## 📊 Benchmark Results

This is the result of running the `./test/benchmark_test.go` benchmark on my machine, (ns/op)times do not include the
terminal graphical visualization time.

| Encoder             | Configuration      | ns/op | B/op | allocs/op |
|---------------------|--------------------|-------|------|-----------|
| **Default Encoder** | All Properties OFF | 490.3 | 80   | 1         |
|                     | All Properties ON  | 511.2 | 104  | 1         |
| **JSON Encoder**    | All Properties OFF | 513.3 | 80   | 1         |
|                     | All Properties ON  | 536.5 | 104  | 1         |
| **YAML Encoder**    | All Properties OFF | 535.3 | 80   | 1         |
|                     | All Properties ON  | 557.1 | 104  | 1         |


## 🤝 Contributing

Contributions are welcome, Here's how you can help:

1. Fork the repository
2. Clone your fork:
3. Create a new branch:

```bash
   git checkout -b feat/your-feature-name
   ```

- **Code Style**
    - Follow standard Go formatting (`go fmt`)
    - Use meaningful variable names
    - Add comments for non-obvious code sections
    - Write tests for new functionality

- **Testing**
    - Run tests: `make test`
    - Run benchmarks: `make test-benchmark`
    - Ensure test coverage remains high; it can be checked using `make test-coverage`

## 📝 License

MIT License—see [LICENSE](https://mit-license.org/) file for details
