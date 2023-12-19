## Tiny Logger

A lightweight logger library implementing standard logging methods with the addition of some colors.    
Logging with this library can be achieved in two main ways:

1. Using the Package Level Logger
2. Instantiating and using a new Logger struct

Both of them implement the same features and methods, it just depends on wether you need
a new Logger instance struct or not.

### Package Logger

The Package logger does not need to be instantiated and is unique when using the library.

```go
logs.Warn("my warning test") // stdout: 'WARN: my warning test'
logs.Info("my", "into", "test", 2) // stdout: 'INFO: my info test 2'
logs.Debug("hey", "check this", "debug") // stdout: 'DEBUG: hey check this debug'
logs.Error("here is the error") // stderr: 'ERROR: here is the error'
```

### Struct Logger

To instantiate a new Logger struct you just need to call the `logs.NewLogger()` method, it will return a pointer
to the newly created Logger.

```go
logger := logs.NewLogger()
logger.Warn("my warning test") // stdout: 'WARN: my warning test'
logger.Info("my", "into", "test", 2) // stdout: 'INFO: my info test 2'
logger.Debug("hey", "check this", "debug") // stdout: 'DEBUG: hey check this debug'
logger.Error("here is the error") // stderr: 'ERROR: here is the error'
```

### Log Levels

Every Logger (Package one included) will have a Log Level property reference.    
This property will tell the Loggers which message categories we want to actually print, based on an internal
hierarchy.    
The logging levels hierarchy is organized as follows:

```go
Error: 0
Warn:  1
Info:  2
Debug: 3
```

Every Logger when firstly generated will have the Log Level set to `Debug` by default, therefore all the logs will be
printed to the
output since `Debug` is the highest in the hierarchy.    
Setting the Logger Log Level to a lower value, let's say `Warn` will cause the Logger to print only `Warn` and `Error`
messages
excluding the categories that are higher of the set Log Level, `Info` and `Debug` for this example.

#### The Log Level can be set in two ways:

1. Using the method
    - Package Logger: `logs.SetLogLvl(WarnLvlName)`
    - Logger Struct: `logs.NewLogger().SetLogLvl(WarnLvlName)`
2. Setting an ENV variable reference. In this case the Logger will retrieve the Log Level from the specified environment
   variable when calling the `SetLogLvlEnvVariable()` method.
    - Package Logger: `SetLogLvlEnvVariable("MY_LOGLVL_ENV_VAR_NAME")`
    - Logger Struct: `logs.NewLogger().SetLogLvlEnvVariable("MY_LOGLVL_ENV_VAR_NAME")`
