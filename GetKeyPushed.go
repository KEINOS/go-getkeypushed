package gkp

import (
	"github.com/mattn/go-tty"
)

// GetKeyPushed returns the single key pressed by the user via tty(terminal).
// It is useful to get the user input without entering `enter` key.
func GetKeyPushed() (string, error) {
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
