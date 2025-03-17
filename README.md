## Tiny Logger

A lightweight logging library with standard logging methods,
supporting customizable log levels, color-coding, and optional date/time formatting.

### Usage

To get started, instantiate a Logger by calling logs.NewLogger(). This method returns a pointer to a new Logger
instance:

```go
logger := logs.NewLogger()
logger.Warn("my warning test") // stdout: 'WARN: my warning test'
logger.Info("my", "into", "test", 2) // stdout: 'INFO: my info test 2'
logger.Debug("hey", "check this", "debug") // stdout: 'DEBUG: hey check this debug'
logger.Error("here is the error") // stderr: 'ERROR: here is the error'
```

The Logger struct implements the [Builder design pattern](https://refactoring.guru/design-patterns/builder) allowing you
to dynamically configure various settings.

### Configuration Options

Use builder methods to customize logging behavior, such as setting the log level, enabling colors, or adding date/time
stamps.

```go
logger := logs.NewLogger().
   SetLogLvl(ll.WarnLvlName).
   EnableColors(true).
   AddTime(true).AddDate(true)

logger.Warn("This is my Warn log", "Test arg")
// OUTPUT: WARN[03/11/2024 18:35:43]: This is my Warn log Test arg

logger.SetEncoder(shared.JsonEncoderType)
logger.Debug("This is my Debug log", "Test arg")
// OUTPUT: {"level":"DEBUG","date":"03/11/2024","time":"18:35:43","message":"This is my Debug log Test arg"}

logger.AddTime(false)
logger.Debug("This is my Debug log", "Test arg")
// OUTPUT: {"level":"DEBUG","date":"03/11/2024","message":"This is my Debug log Test second arg"}
```

### Log Levels

Each `Logger` has a `Log Level` property that determines which message categories to print, based on a priority
hierarchy.
The log levels are ordered as follows:

```go
Error: 0
Warn:  1
Info:  2
Debug: 3
```

By default, the `Logger` is set to the `Debug` level, so all log levels will be printed to the output.
Adjusting the `Log Level` to a lower setting, such as `Warn`, will limit output to only `Warn` and `Error` messages,
filtering out `Info` and `Debug`.

### The Log Level can be set in two ways:

1. Using the method

```go
logs.NewLogger().SetLogLvl(WarnLvlName)
```

2. Using an Environment Variable: Set the log level through an environment variable and configure the Logger to retrieve
   it using the `SetLogLvlEnvVariable()` method.

```go
logs.NewLogger().SetLogLvlEnvVariable("MY_LOGLVL_ENV_VAR_NAME")
```
