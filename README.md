# Tiny Logger

A fast, lightweight, zero-dependency logging solution for Go applications that prioritizes performance and simplicity.  

The library is extremely optimized to log loosely-typed data (interfaces), so you won't have to specify concrete types before logging.  

I know that higher raw speed can be reached using other Go logging solutions, but when I created tiny-logger, I wanted to build something as fast as possible without compromising on simplicity of use.  
There are many projects that can benefit from having a logging library that is compact, fast, easy to use, and simple to modify.  
Since the codebase is so small, it won't take long for you to understand it.

The project is compatible with **Go version 1.18.x** and above.

## Key Features

- **Lightweight**: The library has no dependencies, the code you see is all that runs.  
    NOTE: The only dependencies you'll see in the `go.mod` file are not included in the final binary since they are only used in `_test` files.
- **Simplicity**: I designed the API to have a minimal learning curve. You'll set it up in seconds.
- **Performance**: The library is benchmarked to be very fast. It implements custom JSON and YAML marshaling
  specifically optimized for logging
  - Up to 1.4x faster JSON marshaling than `encoding/json`
  - Up to 5x faster YAML marshaling than `gopkg.in/yaml.v3`
- **Color Support**: Built-in ANSI color support for terminal output
- **Thread-Safe**: Concurrent-safe logging with atomic operations
- **Time-Optimized**: Efficient date/time formatting with minimal allocations

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

## Benchmark Results

1. Benchmarks of the **loggers-comparison-benchmark** command contained in the [makefile](./makefile)

    - **OS:** Linux
    - **Arch:** AMD64
    - **CPU:** 12th Gen Intel(R) Core(TM) i9-12900K

    | Logger | Iterations | Time / Op | Bytes / Op | Allocs / Op |
    | :--- | :--- | :--- | :--- | :--- |
    | **TinyLogger** | 17,625,723 | **339.9 ns** | **88 B** | **2** |
    | Zerolog | 12,983,034 | 460.2 ns | 232 B | 5 |
    | Zap | 10,391,967 | 578.3 ns | 136 B | 2 |
    | Logrus | 3,607,248 | 1692 ns | 1241 B | 21 |

    - **OS:** Darwin (macOS)
    - **Arch:** AMD64
    - **CPU:** VirtualApple @ 2.50GHz

    | Logger | Iterations | Time / Op | Bytes / Op | Allocs / Op |
    | :--- | :--- | :--- | :--- | :--- |
    | **TinyLogger** | 6,091,185 | **972.9 ns** | **88 B** | **2** |
    | Zerolog | 4,922,115 | 1220 ns | 232 B | 5 |
    | Zap | 3,938,301 | 1517 ns | 136 B | 2 |
    | Logrus | 1,814,809 | 3291 ns | 1241 B | 21 |

2. Benchmarks of the **encoders-benchmark** command contained in the [makefile](./makefile)

    - **OS:** Linux
    - **Arch:** AMD64
    - **CPU:** 12th Gen Intel(R) Core(TM) i9-12900K

    | Logger | Iterations | Time / Op | Bytes / Op | Allocs / Op |
    | :--- | :--- | :--- | :--- | :--- |
    | DefaultEncoder DisabledProperties | 18336217 | 298.7 ns | 88 B | 2 |
    | DefaultEncoder EnabledProperties | 18336217 | 334.3 ns | 88 B | 2 |
    | JsonEncoder DisabledProperties | 17974824 | 316.0 ns | 88 B | 2 |
    | JsonEncoder EnabledProperties | 17488896 | 344.2 ns | 88 B | 2 |
    | YamlEncoder DisabledProperties | 17625220 | 342.8 ns | 88 B | 2 |
    | YamlEncoder EnabledProperties | 16005187 | 373.3 ns | 88 B | 2 |

## Contributing

Contributions to this project are really welcome, here's how you can do it:

  1. Fork the repository
  2. Clone your fork
  3. Create a new branch
    ```bash git checkout -b your-feature-name```
  4. Local Tests  
    Take a look at the [makefile](./makefile).  
    You can use the contained commands to run `unit-test`, check the `testing coverage` and monitor the `benchamrks` of the library.

## License

MIT License—see [LICENSE](https://mit-license.org/) file for details  
