package logger

import (
	"fmt"
	"io"
	"log"
	"os"
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

func init() {
	l = New()
}

type Logger struct {
	T, D, I, W, E *log.Logger
	t, d, i, w, e bool
}

func New() (l *Logger) {
	l = &Logger{
		T: log.New(os.Stdout, "[trace] ", log.Lmicroseconds|log.Lshortfile),
		D: log.New(os.Stdout, "[debug] ", log.Ltime|log.Lshortfile),
		I: log.New(os.Stdout, "", log.Ldate|log.Ltime),
		W: log.New(os.Stdout, "[warning] ", log.Ldate|log.Ltime),
		E: log.New(os.Stdout, "[error] ", log.Ldate|log.Ltime|log.Lshortfile),
		t: true,
		d: true,
		i: true,
		w: true,
		e: true,
	}
	return
}

func SetOutput(w io.Writer) {
	l.SetOutput(w)
}

func SetLevel(s string) {
	l.SetLevel(s)
}

func (l *Logger) SetOutput(w io.Writer) {
	l.T.SetOutput(w)
	l.D.SetOutput(w)
	l.I.SetOutput(w)
	l.W.SetOutput(w)
	l.E.SetOutput(w)
}

func (l *Logger) SetLevel(s string) {
	l.t = true
	l.d = true
	l.i = true
	l.w = true
	l.e = true
	switch s {
	case "debug":
		l.t = false
	case "info":
		l.t = false
		l.d = false
	case "warn":
		l.t = false
		l.d = false
		l.i = false
	case "error":
		l.t = false
		l.d = false
		l.i = false
		l.w = false
	}
}

func Trace(format string, v ...interface{}) {
	l.Trace(format, v...)
}

func Debug(format string, v ...interface{}) {
	l.Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	l.Info(format, v...)
}

func Warn(format string, v ...interface{}) {
	l.Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	l.Error(format, v...)
}

func (l *Logger) Trace(format string, v ...interface{}) {
	if l.t {
		l.T.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Debug(format string, v ...interface{}) {
	if l.d {
		l.D.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	if l.i {
		l.I.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Warn(format string, v ...interface{}) {
	if l.w {
		l.W.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Error(format string, v ...interface{}) {
	if l.e {
		l.E.Output(2, fmt.Sprintf(format, v...))
	}
}
