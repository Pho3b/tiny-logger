package logs

import (
	"fmt"
	"os"
	"strings"
)

// Log Level STRING Constants
const (
	ErrorLvlName = "ERROR"
	WarnLvlName  = "WARN"
	InfoLvlName  = "INFO"
	DebugLvlName = "DEBUG"
)

// Log Level INT8 Constants
const (
	ErrorLvl = int8(0)
	WarnLvl  = int8(1)
	InfoLvl  = int8(2)
	DebugLvl = int8(3)
)

// Log Level INT to STRING map
var logLvlIntToName = map[int8]string{
	ErrorLvl: ErrorLvlName,
	WarnLvl:  WarnLvlName,
	InfoLvl:  InfoLvlName,
	DebugLvl: DebugLvlName,
}

// Log Level STRING to INT map
var logLvlNameToInt = map[string]int8{
	ErrorLvlName: ErrorLvl,
	WarnLvlName:  WarnLvl,
	InfoLvlName:  InfoLvl,
	DebugLvlName: DebugLvl,
}

// debug prints the given objects as strings to the 'standard output' and colors the prefix in GREY if
// supported by the operating system.
func debug(args ...interface{}) {
	if len(args) > 0 {
		_, _ = fmt.Fprintln(os.Stdout, fmt.Sprintf("%vDEBUG:%v %s", gray, reset, buildMsg(args...)))
	}
}

// info prints the given objects as strings to the 'standard output' and colors the prefix in CYAN in
// supported by the operating system.
func info(args ...interface{}) {
	if len(args) > 0 {
		_, _ = fmt.Fprintln(os.Stdout, fmt.Sprintf("%vINFO:%v %s", cyan, reset, buildMsg(args...)))
	}
}

// warn prints the given objects as strings to the 'standard output' and colors the prefix in YELLOW if
// supported by the operating system.
func warn(args ...interface{}) {
	if len(args) > 0 {
		_, _ = fmt.Fprintln(os.Stdout, fmt.Sprintf("%vWARNING:%v %s", yellow, reset, buildMsg(args...)))
	}
}

// error prints the given objects as strings to the 'standard error' and colors the prefix in RED if
// supported by the operating system.
// It does not print anything if all the given args result to be nil.
func error(args ...interface{}) {
	if len(args) > 0 && !areAllNil(args...) {
		_, _ = fmt.Fprintln(os.Stderr, fmt.Sprintf("%vERROR:%v %s", red, reset, buildMsg(args...)))
	}
}

// fatalError prints the given objects as strings to the 'standard error' and colors the prefix in MAGENTA if
// supported by the operating system, it also exits the current process.
// fatalError does not print anything and does not exit the current process if all the given args result to be nil.
func fatalError(args ...interface{}) {
	if len(args) > 0 && !areAllNil(args...) {
		_, _ = fmt.Fprintln(os.Stderr, fmt.Sprintf("%vFATAL ERROR:%v %s", lightMagenta, reset, buildMsg(args...)))
		os.Exit(1)
	}
}

// buildMsg returns a string containing all the given arguments cast to strings concatenated with a white space.
func buildMsg(args ...interface{}) string {
	res := strings.Builder{}

	for _, arg := range args {
		res.WriteString(fmt.Sprintf("%v ", arg))
	}

	return strings.TrimSuffix(res.String(), " ")
}

// areAllNil returns true if all the given args are 'nil', false otherwise.
func areAllNil(args ...interface{}) bool {
	for _, arg := range args {
		if arg != nil {
			return false
		}
	}

	return true
}

// retrieveLogLvlFromEnv attempts to retrieve the given 'logLvlEnvVariable' value from the ENV variables.
// If the given variable is not found, 'DebugLvl' is returned by default.
func retrieveLogLvlFromEnv(logLvlEnvVariable string) int8 {
	return retrieveLogLvlIntFromName(
		strings.ToUpper(os.Getenv(logLvlEnvVariable)),
	)
}

// retrieveLogLvlIntFromName given a logLvlName returns its int8 constant value.
// If the given logLvlName is not valid, DebugLvl is returned by default.
func retrieveLogLvlIntFromName(logLvlName string) int8 {
	if _, found := logLvlNameToInt[logLvlName]; !found {
		logLvlName = DebugLvlName
	}

	return logLvlNameToInt[logLvlName]
}
