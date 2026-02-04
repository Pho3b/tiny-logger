# Tiny Logger

A fast, lightweight, zero-dependency logging solution for Go applications that prioritizes performance and simplicity.  

The library is extremely optimized to log loosely typed data (interfaces), so you won't have to specify the correct data type before logging them.  
I know you can reach more raw speed using other Go logging libray and specifying the type of the data to log beforehead, but when I created the tiny-logger i wanted to make it as fast as possible while keeping it very compact and simple to use.  
In my experience there are cases where the only thing you know is raw speed yes, but majority of projects benefits from having something that is really small, easy to use and also perform very fast.  

Compatible with Go version 1.18.x and above

## Key Features

- **Lightweight**: The library has no dependencies, the code you see is all that runs.  
  NOTE: The only dpes you'll see in the `go.mod` file are not included in the final binary since they are only used in `_test` files.
- **Simplicity**: I made the API to have a minimal learning curve. You'll set it up in seconds.
- **Performance**: The library is benchmarked to be very fast. It implements custom JSON and YAML marshaling
  specifically optimized for logging
  - Up to 1.4x faster JSON marshaling than `encoding/json`
  - Up to 5x faster YAML marshaling than `gopkg.in/yaml.v3`
- **Color Support**: Built-in ANSI color support for terminal output
- **Thread-Safe**: Concurrent-safe logging with atomic operations
- **Time-Optimized**: Efficient date/time print built-int logic with minimal allocations

## Use Examples

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

## Benchmark Results

**System Environment:**

- **OS:** Linux
- **Arch:** AMD64
- **CPU:** 12th Gen Intel(R) Core(TM) i9-12900K

| Logger | Iterations | Time / Op | Bytes / Op | Allocs / Op |
| :--- | :--- | :--- | :--- | :--- |
| **TinyLogger** | 17,625,723 | **339.9 ns** | **88 B** | **2** |
| Zerolog | 12,983,034 | 460.2 ns | 232 B | 5 |
| Zap | 10,391,967 | 578.3 ns | 136 B | 2 |
| Logrus | 3,607,248 | 1692 ns | 1241 B | 21 |

## Contributing

Contributes to this project are really welcome, here's how you can do it

  1. Fork the repository
  2. Clone your fork
  3. Create a new branch
    ```bash git checkout -b your-feature-name```
  4. Testing  
    Take a look at the [makefile](./makefile).  
    You can use the contained commands to run `unit-test`, check the `testing coverage` and monitor the `benchamrks` of the library.

## 📝 License

MIT License—see [LICENSE](https://mit-license.org/) file for details  
