package funcs_test

import (
	"errors"
	"testing"

	"github.com/KEINOS/go-getkeypushed/funcs"
	"github.com/mattn/go-tty"
	"github.com/stretchr/testify/assert"
)

func TestOpenTTY(t *testing.T) {
	// Backup and restore the original function
	OldTTYOpen := funcs.TTYOpen
	defer func() {
		funcs.TTYOpen = OldTTYOpen
	}()

	// Mock function
	funcs.TTYOpen = func() (*tty.TTY, error) {
		return nil, errors.New("foo")
	}

	// Test
	result, err := funcs.OpenTTY()

	assert.Nil(t, result, "when fails to open TTY should return an error")
	assert.Error(t, err)
}

func TestOpenTTY_mock_to_dummy_tty(t *testing.T) {
	// Backup and restore the original function
	OldTTYOpen := funcs.TTYOpen
	defer func() {
		funcs.TTYOpen = OldTTYOpen
	}()

	// Mock function
	funcs.TTYOpen = func() (*tty.TTY, error) {
		return new(tty.TTY), nil
	}

	// Test
	result, err := funcs.OpenTTY()

	assert.Nil(t, err, "testing does not have a TTY")
	assert.IsType(t, new(tty.TTY), result)
}
