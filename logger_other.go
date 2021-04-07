// +build !js

package logger

import "fmt"

func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.traceEnabled {
		l.TraceLogger.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.debugEnable {
		l.DebugLogger.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.infoEnabled {
		l.InfoLogger.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.warnEnabled {
		l.WarnLogger.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.errorEnabled {
		l.ErrorLogger.Output(3, fmt.Sprintf(format, v...))
	}
}
