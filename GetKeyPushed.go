package gkp

import (
	"github.com/mattn/go-tty"
)

var (
	// DummyMsg is a string that GetKeyPushed() returns if set. Use this for testing.
	DummyMsg string = ""
	// DummyErr is an error message that GetKeyPushed() returns if set. Use this for testing.
	DummyErr error = nil
)

// GetKeyPushed returns the single key pressed by the user via tty(terminal).
// It is useful to get the user input without entering `enter` key.
func GetKeyPushed() (string, error) {
	// Return dummy if set. It will be used to ease tesing at the other
	// packages that uses this pakage.
	if DummyMsg != "" || DummyErr != nil {
		return DummyMsg, DummyErr
	}

	var (
		t   *tty.TTY
		r   rune
		err error
	)

	if t, err = tty.Open(); err != nil {
		return "", err
	}

	defer t.Close()

	if r, err = t.ReadRune(); err != nil {
		return "", err
	}

	return string(r), nil
}
