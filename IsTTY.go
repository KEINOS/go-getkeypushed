package gkp

import (
	"os"

	"github.com/mattn/go-isatty"
)

// DummyIsTTY is a bool that IsTTY() returns if set. Use this for testing.
var DummyIsTTY bool = false

// IsTTY returns false if the program is not running from terminal.
func IsTTY() bool {
	if isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd()) || DummyIsTTY {
		return true
	}

	return false
}
