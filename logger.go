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
		T: log.New(os.Stdout, "[trace]\t", log.Lmicroseconds|log.Lshortfile),
		D: log.New(os.Stdout, "[debug]\t", log.Ltime|log.Lshortfile),
		I: log.New(os.Stdout, "[info]\t", log.Ldate|log.Ltime),
		W: log.New(os.Stdout, "[warn]\t", log.Ldate|log.Ltime),
		E: log.New(os.Stdout, "[error]\t", log.Ldate|log.Ltime|log.Lshortfile),
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

func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.t {
		l.T.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.d {
		l.D.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.i {
		l.I.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.w {
		l.W.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.e {
		l.E.Output(2, fmt.Sprintf(format, v...))
	}
}
