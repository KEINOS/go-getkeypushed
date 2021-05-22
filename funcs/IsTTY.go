package funcs

import (
	"os"

	"github.com/mattn/go-isatty"
)

// DummyIsTTY is a bool that IsTTY() returns if set. Use this for testing.
var DummyIsTTY bool = false

// IsTTY returns false if the program is not running from terminal.
//
// To mock its behavior tor testing, set funcs.DummyIsTTY to true.
func IsTTY() bool {
	if isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd()) || DummyIsTTY {
		return true
	}

	return false
}
