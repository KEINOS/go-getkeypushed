package getkbd

import (
	"errors"

	"github.com/KEINOS/go-getkeypushed/funcs"
)

var (
	// DummyKey is a string that GetKeyPushed() returns if set. Use this for testing.
	DummyKey string = ""
	// DummyErr is an error message that GetKeyPushed() returns if set. Use this for testing.
	DummyErr   error = nil
	DummyIsTTY bool  = false
	// UserInputDummy is a string that GetKeyPushed() returns if set. Use this for testing.
	UserInputDummy string = ""
)

// getkbd.Pushed returns a single char pushed by the user via tty(terminal).
//
// It is useful to get the user input without entering `enter` key.
func Pushed(keyDefault string, timeWait int) (keyPushed string, err error) {
	// Return the dummy value if set.
	// It will be used to ease tesing at the other packages that uses this pakage.
	if DummyKey != "" || DummyErr != nil {
		return DummyKey, DummyErr
	}

	// Pass through the value to mock funcs.GetInput()
	funcs.UserInputDummy = UserInputDummy
	if DummyIsTTY {
		funcs.DummyIsTTY = DummyIsTTY
	}

	if !funcs.IsTTY() {
		return "", errors.New("not a TTY")
	}

	timer := funcs.StartTimer(timeWait)
	defer func() { timer.Stop() }()

	return funcs.GetInput(timer, keyDefault)
}
