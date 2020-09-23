# LiteLog

## Litelog is simple tiny package that provide logging with levels
There are five levels 
```go
// LOG levels
const (
	Info = iota
	Warn
	Err
	Debug
	Trace
)
```
If defined level is less than called method expected, than message will not be logged.
It provide you ability to increase verbosity of logs.

Example

```go
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
1. Time prefix with ```go litelog.WithTime(layout ...string)``` method. Layout is a some time package layout const
layout parameter can be empty, than ```go litelog.WithTime()``` method will set time.RFC3339 layout
2. Custom text prefix before message with ```go litelog.WithPrefix(pr string)``` method.
3. Custom io.Writer object where Logger will be write mesages with ```go litelog.WithWriter(wr io.Writer)``` method (By default os.Stdout).
4. Log level with ```go litelog.WithLevel(lev int)``` method.

# Level dependent methods
* ```go Info(msg string)``` prints message if current level is more or equal Info with "[INFO]" prefix
* ```go Infof(msg string, args ...interface{})``` prints formatted message if current level is more or equal Info with "[INFO]" prefix
* ```go Warn(msg string)``` prints message if current level is more or equal Warn with "[WARN]" prefix
* ```go Warnf(msg string, args ...interface{})``` prints formatted message if current level is more or equal Warn with "[WARN]" prefix
* ```go Error(msg string)``` prints message if current level is more or equal Err with "[ERROR]" prefix
* ```go Errorf(msg string, args ...interface{})``` prints formatted message if current level is more or equal Err with "[ERROR]" prefix
* ```go Debug(msg string)``` prints message if current level is more or equal Debug with "[DEBUG]" prefix
* ```go Debugf(msg string, args ...interface{})``` prints formatted message if current level is more or equal Debug with "[DEBUG]" prefix
* ```go Trace(msg string)``` prints message if current level is more or equal Trace with "[TRACE]" prefix
* ```go Tracef(msg string, args ...interface{})``` prints formatted message if current level is equal Trace with "[TRACE]" prefix

# Level independent methods
* ```go Println(msg string)``` prints message in provided io.Writer
* ```go Printf(msg string, args ...interface{})``` prints formatted message in provided io.Writer
* ```go Fatal(msg string)``` prints message and exit with status code 1
* ```go Fatalf(msg string, args ...interface{})``` prints formatted message and exit with status code 1