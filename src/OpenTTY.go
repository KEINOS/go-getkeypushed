package gkp

import (
	"fmt"

	"github.com/mattn/go-tty"
)

// TTYOpen is a copy of tty.Open() to ease testing. Override the function for testing.
var TTYOpen func() (*tty.TTY, error) = tty.Open

func OpenTTY() (t *tty.TTY, err error) {
	if t, err = TTYOpen(); err != nil {
		return nil, fmt.Errorf("failed to open TTY in OpenTTY().\nErrMsg: %v", err)
	}

	return t, nil
}
