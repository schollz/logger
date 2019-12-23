package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestTrace(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)

	SetLevel("trace")
	funcs := []func(...interface{}){Trace, Debug, Info, Warn, Error}
	for i, s := range []string{"trace", "debug", "info", "warn", "error"} {
		funcs[i](s)
		fmt.Print(buf.String())
		if !strings.Contains(buf.String(), s) {
			t.Error("bad " + s)
		}
		buf.Reset()
	}
}

func TestLogging(t *testing.T) {
	logger := New()
	var buf bytes.Buffer
	logger.SetOutput(&buf)

	levels := []string{"trace", "debug", "info", "warn", "error"}
	funcs := []func(string, ...interface{}){logger.Tracef, logger.Debugf, logger.Infof, logger.Warnf, logger.Errorf}

	for level, levelName := range levels {
		logger.SetLevel(levelName)
		for i, s := range levels {
			funcs[i](s)
			if i < level {
				if buf.String() != "" {
					t.Error("bad " + s)
				}
			} else {
				if !strings.Contains(buf.String(), s) {
					t.Error("bad " + s)
				}
			}
			buf.Reset()
			if logger.GetLevel() != levelName {
				t.Error("could not get level name: " + levelName)
			}
		}
	}
}
