package gkp

import (
	"errors"
	"os"
)

var (
	// DummyKey is a string that GetKeyPushed() returns if set. Use this for testing.
	DummyKey string = ""
	// DummyErr is an error message that GetKeyPushed() returns if set. Use this for testing.
	DummyErr error = nil
	// UserInput is a string that GetKeyPushed() returns if set. Use this for testing.
	// It will override the user input.
	UserInput string = ""
	// OsExit is a copy of os.Exit to ease mock the application exit.
	OsExit func(code int) = os.Exit
)

// GetKeyPushed returns the single key pressed by the user via tty(terminal).
//
// It is useful to get the user input without entering `enter` key.
func GetKeyPushed(keyDefault string, timeWait int) (keyPushed string, err error) {
	// Return the dummy value if set.
	// It will be used to ease tesing at the other packages that uses this pakage.
	if DummyKey != "" || DummyErr != nil {
		return DummyKey, DummyErr
	}

	if !IsTTY() {
		return "", errors.New("not a TTY")
	}

	timer := startTimer(timeWait)
	defer timer.Stop()

	if timeWait == 0 {
		if !timer.Stop() {
			<-timer.C
		}
	}

	return getInput(timer, keyDefault)
}
