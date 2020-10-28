package log

import "fmt"

type loglevel int

var tracell loglevel
var debugll loglevel
var infoll loglevel
var warnll loglevel
var errorll loglevel
var fatalll loglevel
var currentll loglevel

func init() {
	tracell = 0
	debugll = 10
	infoll = 20
	warnll = 30
	errorll = 40
	fatalll = 50
	currentll = tracell
}

func Trace(msg string) {
	if currentll <= tracell {
		logMessage("TRACE", msg)
	}
}

func Debug(msg string) {
	if currentll <= debugll {
		logMessage("DEBUG", msg)
	}
}

func Info(msg string) {
	if currentll <= infoll {
		logMessage("INFO", msg)
	}
}

func Warn(msg string) {
	if currentll <= warnll {
		logMessage("WARN", msg)
	}
}

func Error(err error, msg string) {
	if currentll <= errorll {
		logMessageWithError("ERROR", err, msg)
	}
}

func Fatal(err error, msg string) {
	if currentll <= fatalll {
		logMessageWithError("FATAL", err, msg)
	}
}

func Start(file string, function string) {
	log(file, function, "start")
}

func End(file string, function string) {
	log(file, function, "end")
}

func log(file string, function string, startend string) {
	fmt.Println("---", file+".go", startend, function, "---")
}

func logMessage(lvl string, msg string) {
	fmt.Println(lvl+":", msg)
}

func logMessageWithError(lvl string, err error, msg string) {
	fmt.Println(lvl+":", err, msg)
}
