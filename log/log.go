package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
)

var stdout io.Writer = color.Output
var g_rl *readline.Instance = nil
var debug_output = true
var mtx_log *sync.Mutex = &sync.Mutex{}

const (
	DEBUG = iota
	INFO
	IMPORTANT
	WARNING
	ERROR
	FATAL
	SUCCESS
)

var LogLabels = map[int]string{
	DEBUG:     "dbg",
	INFO:      "inf",
	IMPORTANT: "imp",
	WARNING:   "war",
	ERROR:     "err",
	FATAL:     "!!!",
	SUCCESS:   "+++",
}

// File logger
var fileLogger *log.Logger
var logFile *os.File
var enableFileLog bool = false

func DebugEnable(enable bool) {
	debug_output = enable
}

func SetOutput(o io.Writer) {
	stdout = o
}

func SetReadline(rl *readline.Instance) {
	g_rl = rl
}

func GetOutput() io.Writer {
	return stdout
}

func NullLogger() *log.Logger {
	return log.New(io.Discard, "", 0)
}

func InitFileLogger(path string, prefix string, flag int) error {
	var err error
	logFile, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	fileLogger = log.New(logFile, prefix, flag)
	enableFileLog = true
	return nil
}

func CloseFileLogger() {
	if logFile != nil {
		logFile.Close()
	}
	enableFileLog = false
}

func refreshReadline() {
	if g_rl != nil {
		g_rl.Refresh()
	}
}

func logToFile(msg string) {
	if enableFileLog && fileLogger != nil {
		fileLogger.Print(msg)
	}
}

func Debug(format string, args ...interface{}) {
	mtx_log.Lock()
	defer mtx_log.Unlock()

	if debug_output {
		msg := format_msg(DEBUG, format+"\n", args...)
		fmt.Fprint(stdout, msg)
		logToFile(msg)
		refreshReadline()
	}
}

func Info(format string, args ...interface{}) {
	mtx_log.Lock()
	defer mtx_log.Unlock()

	msg := format_msg(INFO, format+"\n", args...)
	fmt.Fprint(stdout, msg)
	logToFile(msg)
	refreshReadline()
}

func Important(format string, args ...interface{}) {
	mtx_log.Lock()
	defer mtx_log.Unlock()

	msg := format_msg(IMPORTANT, format+"\n", args...)
	fmt.Fprint(stdout, msg)
	logToFile(msg)
	refreshReadline()
}

func Warning(format string, args ...interface{}) {
	mtx_log.Lock()
	defer mtx_log.Unlock()

	msg := format_msg(WARNING, format+"\n", args...)
	fmt.Fprint(stdout, msg)
	logToFile(msg)
	refreshReadline()
}

func Error(format string, args ...interface{}) {
	mtx_log.Lock()
	defer mtx_log.Unlock()

	msg := format_msg(ERROR, format+"\n", args...)
	fmt.Fprint(stdout, msg)
	logToFile(msg)
	refreshReadline()
}

func Fatal(format string, args ...interface{}) {
	mtx_log.Lock()
	defer mtx_log.Unlock()

	msg := format_msg(FATAL, format+"\n", args...)
	fmt.Fprint(stdout, msg)
	logToFile(msg)
	refreshReadline()
}

func Success(format string, args ...interface{}) {
	mtx_log.Lock()
	defer mtx_log.Unlock()

	msg := format_msg(SUCCESS, format+"\n", args...)
	fmt.Fprint(stdout, msg)
	logToFile(msg)
	refreshReadline()
}

func Printf(format string, args ...interface{}) {
	mtx_log.Lock()
	defer mtx_log.Unlock()

	msg := fmt.Sprintf(format, args...)
	fmt.Fprint(stdout, msg)
	logToFile(msg)
	refreshReadline()
}

func format_msg(lvl int, format string, args ...interface{}) string {
	label := LogLabels[lvl]
	if label == "" {
		label = "???"
	}
	msg := fmt.Sprintf(format, args...)
	return fmt.Sprintf("[%s] %s", label, msg)
}
