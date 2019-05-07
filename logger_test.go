package logger

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogging(t *testing.T) {
	logger := New()
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	logger.Trace("hello")
	fmt.Println(buf.String())
	assert.Equal(t, "", buf.String())
}
