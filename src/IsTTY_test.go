package gkp_test

import (
	"testing"

	gkp "github.com/KEINOS/go-getkeypushed/src"
	"github.com/stretchr/testify/assert"
)

func TestIsTTY(t *testing.T) {
	result := gkp.IsTTY()

	assert.False(t, result, "testing does not have a TTY")
}

func TestIsTTY_mock_tty(t *testing.T) {
	OldOsStdoutFd := gkp.DummyIsTTY
	defer func() {
		gkp.DummyIsTTY = OldOsStdoutFd
	}()

	gkp.DummyIsTTY = true

	result := gkp.IsTTY()

	assert.True(t, result)
}
