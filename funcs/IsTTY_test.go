package funcs_test

import (
	"testing"

	"github.com/KEINOS/go-getkeypushed/funcs"
	"github.com/stretchr/testify/assert"
)

func TestIsTTY(t *testing.T) {
	result := funcs.IsTTY()

	assert.False(t, result, "testing does not have a TTY")
}

func TestIsTTY_mock_tty(t *testing.T) {
	OldOsStdoutFd := funcs.DummyIsTTY
	defer func() {
		funcs.DummyIsTTY = OldOsStdoutFd
	}()

	funcs.DummyIsTTY = true

	result := funcs.IsTTY()

	assert.True(t, result)
}
