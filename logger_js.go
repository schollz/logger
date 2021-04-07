//+build js

package logger

import (
	"fmt"
	"syscall/js"
)

func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.traceEnabled {
		js.Global().Get("console").Call("trace", fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.debugEnable {
		js.Global().Get("console").Call("debug", fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.infoEnabled {
		js.Global().Get("console").Call("info", fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.warnEnabled {
		js.Global().Get("console").Call("warn", fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.errorEnabled {
		js.Global().Get("console").Call("error", fmt.Sprintf(format, v...))
	}
}
