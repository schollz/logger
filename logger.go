package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

var l *Logger

type Logger struct {
	T, D, I, W, E *log.Logger
	t, d, i, w, e bool
}

func New() (l *Logger) {
	l = &Logger{
		T: log.New(os.Stdout, "trace: ", log.Ltime|log.Lshortfile),
		D: log.New(os.Stdout, "debug: ", log.Ltime|log.Lshortfile),
		I: log.New(os.Stdout, "", log.Ldate|log.Ltime),
		W: log.New(os.Stdout, "warn:  ", log.Ltime),
		E: log.New(os.Stdout, "error: ", log.Ltime|log.Lshortfile),
		t: true,
		d: true,
		i: true,
		w: true,
		e: true,
	}
	return
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
	}
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
