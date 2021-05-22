package funcs

import "os"

var (
	// OsExit is a copy of os.Exit to ease mock the application exit.
	OsExit func(code int) = os.Exit
	// UserInputDummy is a string that gkptty.GetInput() returns if set. Use this for testing.
	// It will override the user input.
	UserInputDummy string = ""
)
