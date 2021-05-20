package gkp_test

import (
	"errors"
	"testing"

	gkp "github.com/KEINOS/go-getkeypushed/src"
	"github.com/mattn/go-tty"
	"github.com/stretchr/testify/assert"
)

func TestOpenTTY(t *testing.T) {
	OldTTYOpen := gkp.TTYOpen
	defer func() {
		gkp.TTYOpen = OldTTYOpen
	}()

	gkp.TTYOpen = func() (*tty.TTY, error) {
		return nil, errors.New("foo")
	}

	result, err := gkp.OpenTTY()

	assert.Nil(t, result, "when fails to open TTY should return an error")
	assert.Error(t, err)
}

func TestOpenTTY_mock_to_dummy_tty(t *testing.T) {
	OldTTYOpen := gkp.TTYOpen
	defer func() {
		gkp.TTYOpen = OldTTYOpen
	}()

	gkp.TTYOpen = func() (*tty.TTY, error) {
		return new(tty.TTY), nil
	}

	result, err := gkp.OpenTTY()

	assert.Nil(t, err, "testing does not have a TTY")
	assert.IsType(t, new(tty.TTY), result)
}
