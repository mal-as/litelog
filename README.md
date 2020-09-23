# LiteLog
<a href="https://codecov.io/gh/mal-as/litelog" alt="coverage">
        <img src="https://img.shields.io/badge/coverage-96.94%25-success" /></a>

## Litelog is simple tiny package that provide logging with levels
There are five levels 
```
// LOG levels
const (
	Info = iota
	Warn
	Err
	Debug
	Trace
)
```
If defined level is less than level that called method is expecting, than message will not be logged.
It provide you ability to increase verbosity of logs.

Example

```
log := litelog.New(
    litelog.WithWriter(os.Stderr) // default os.Stdout
    litelog.WithLevel(litelog.Warn) // default Info
)

log.Error("some error message") // will not be logged
log.Debug("some debug message") // will not be logged

log.Info("some info message") // will be logged ("[INFO] some info message")
log.Warn("some warn message") // will be logged ("[WARN] some info message")
```
# Logger parameters
Package provide ability to add: 
1. Time prefix with ```litelog.WithTime(layout ...string)``` method. Layout is a some time package layout const
layout parameter can be empty, than ```litelog.WithTime()``` method will set ```time.RFC3339``` layout
2. Custom text prefix before message with ```litelog.WithPrefix(pr string)``` method.
3. Custom io.Writer object where Logger will be write mesages with ```litelog.WithWriter(wr io.Writer)``` method (By default os.Stdout).
4. Log level with ```litelog.WithLevel(lev int)``` method.

# Level dependent methods
* ```Info(msg string)``` prints message if current level is more or equal Info with "[INFO]" prefix
* ```Infof(msg string, args ...interface{})``` prints formatted message if current level is more or equal Info with "[INFO]" prefix
* ```Warn(msg string)``` prints message if current level is more or equal Warn with "[WARN]" prefix
* ```Warnf(msg string, args ...interface{})``` prints formatted message if current level is more or equal Warn with "[WARN]" prefix
* ```Error(msg string)``` prints message if current level is more or equal Err with "[ERROR]" prefix
* ```Errorf(msg string, args ...interface{})``` prints formatted message if current level is more or equal Err with "[ERROR]" prefix
* ```Debug(msg string)``` prints message if current level is more or equal Debug with "[DEBUG]" prefix
* ```Debugf(msg string, args ...interface{})``` prints formatted message if current level is more or equal Debug with "[DEBUG]" prefix
* ```Trace(msg string)``` prints message if current level is more or equal Trace with "[TRACE]" prefix
* ```Tracef(msg string, args ...interface{})``` prints formatted message if current level is equal Trace with "[TRACE]" prefix

# Level independent methods
* ```Println(msg string)``` prints message in provided io.Writer
* ```Printf(msg string, args ...interface{})``` prints formatted message in provided io.Writer
* ```Fatal(msg string)``` prints message and exit with status code 1
* ```Fatalf(msg string, args ...interface{})``` prints formatted message and exit with status code 1
