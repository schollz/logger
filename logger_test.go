package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestTrace(t *testing.T) {
	logger := New()
	var buf bytes.Buffer
	logger.SetOutput(&buf)

	logger.SetLevel("trace")
	funcs := []func(string, ...interface{}){logger.Trace, logger.Debug, logger.Info, logger.Warn, logger.Error}
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
	funcs := []func(string, ...interface{}){logger.Trace, logger.Debug, logger.Info, logger.Warn, logger.Error}

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
		}
	}

}
