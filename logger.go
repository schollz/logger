package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

var l *Logger

type writer struct {
	io.Writer
	timeFormat string
}

func (w writer) Write(b []byte) (n int, err error) {
	return w.Writer.Write(append([]byte(time.Now().Format(w.timeFormat)), b...))
}

const (
	red    = "\033[0;31;1m"
	yellow = "\033[0;33m"
	white  = "\033[0;37m"
	cyan   = "\033[0;36m"
	blue   = "\033[0;34;1m"
	end    = "\033[0m"
)

func init() {
	l = New()
}

type Logger struct {
	TraceLogger, DebugLogger, InfoLogger, WarnLogger, ErrorLogger     *log.Logger
	traceEnabled, debugEnable, infoEnabled, warnEnabled, errorEnabled bool
}

func New() (l *Logger) {
	l = &Logger{
		TraceLogger: log.New(os.Stdout, "[trace]\t", log.Ltime|log.Lmicroseconds|log.Lshortfile),
		DebugLogger: log.New(os.Stdout, "[debug]\t", log.Ltime|log.Lshortfile),
		InfoLogger:  log.New(os.Stdout, "[info]\t", log.Ldate|log.Ltime),
		WarnLogger:  log.New(os.Stdout, "[warn]\t", log.Ldate|log.Ltime),
		ErrorLogger: log.New(os.Stdout, "[error]\t", log.Ldate|log.Ltime|log.Lshortfile),
		traceEnabled: true,
		debugEnable:  true,
		infoEnabled:  true,
		warnEnabled:  true,
		errorEnabled: true,
	}
	if runtime.GOOS == "linux" {
		l.TraceLogger.SetPrefix(blue + l.TraceLogger.Prefix() + end)
		l.DebugLogger.SetPrefix(cyan + l.DebugLogger.Prefix() + end)
		l.InfoLogger.SetPrefix(white + l.InfoLogger.Prefix() + end)
		l.WarnLogger.SetPrefix(yellow + l.WarnLogger.Prefix() + end)
		l.ErrorLogger.SetPrefix(red + l.ErrorLogger.Prefix() + end)
	}
	if strings.TrimSpace(strings.ToLower(os.Getenv("LOGGER"))) != "" {
		l.SetLevel(strings.TrimSpace(strings.ToLower(os.Getenv("LOGGER"))))
	}
	return
}

func SetOutput(w io.Writer) {
	l.SetOutput(w)
}

func SetLevel(s string) {
	// LOGGER enviromental variable takes precedence
	if strings.TrimSpace(strings.ToLower(os.Getenv("LOGGER"))) != "" {
		return
	}
	l.SetLevel(s)
}

func (l *Logger) SetOutput(w io.Writer) {
	l.TraceLogger.SetOutput(w)
	l.DebugLogger.SetOutput(w)
	l.InfoLogger.SetOutput(w)
	l.WarnLogger.SetOutput(w)
	l.ErrorLogger.SetOutput(w)
}

func (l *Logger) SetLevel(s string) {
	l.traceEnabled = true
	l.debugEnable = true
	l.infoEnabled = true
	l.warnEnabled = true
	l.errorEnabled = true
	switch s {
	case "debug":
		l.traceEnabled = false
	case "info":
		l.traceEnabled = false
		l.debugEnable = false
	case "warn":
		l.traceEnabled = false
		l.debugEnable = false
		l.infoEnabled = false
	case "error":
		l.traceEnabled = false
		l.debugEnable = false
		l.infoEnabled = false
		l.warnEnabled = false
	}
}

func GetLevel() (s string) {
	return l.GetLevel()
}

func (l *Logger) GetLevel() (s string) {
	if l.traceEnabled {
		return "trace"
	} else if l.debugEnable {
		return "debug"
	} else if l.infoEnabled {
		return "info"
	} else if l.warnEnabled {
		return "warn"
	}
	return "error"
}

func Tracef(format string, v ...interface{}) {
	l.Tracef(format, v...)
}

func Debugf(format string, v ...interface{}) {
	l.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	l.Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	l.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	l.Errorf(format, v...)
}

func Trace(v ...interface{}) {
	l.Tracef(fmt.Sprint(v...))
}

func Debug(v ...interface{}) {
	l.Debugf(fmt.Sprint(v...))
}

func Info(v ...interface{}) {
	l.Infof(fmt.Sprint(v...))
}

func Warn(v ...interface{}) {
	l.Warnf(fmt.Sprint(v...))
}

func Error(v ...interface{}) {
	l.Errorf(fmt.Sprint(v...))
}