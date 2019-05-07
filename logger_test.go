package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestLogging(t *testing.T) {
	logger := New()
	var buf bytes.Buffer
	logger.SetOutput(&buf)

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
