package getkbd_test //nolint:testpackage // To override osExit the test needs to be part of main

import (
	"errors"
	"testing"

	getkbd "github.com/KEINOS/go-getkeypushed"

	"github.com/stretchr/testify/assert"
)

func TestPushed_DummyErr(t *testing.T) {
	expect := "dummy error occurred"

	errDummy := errors.New(expect)

	getkbd.DummyErr = errDummy

	defer func() { getkbd.DummyErr = nil }()

	msg, err := getkbd.Pushed("q", 5)

	assert.Equal(t, "", msg,
		"When DummyErr global variable is not a nil then Pushed() should be empty('') message.")

	assert.NotNil(t, err,
		"When DummyErr global variable is not a nil then Pushed() should return an error.")

	assert.EqualErrorf(t, err, expect, "error message %s", "mal-formatted")
}

func TestPushed_DummyKey(t *testing.T) {
	expect := "a"

	getkbd.DummyKey = expect

	defer func() { getkbd.DummyKey = "" }()

	actual, err := getkbd.Pushed("q", 5)
	assert.Equal(t, expect, actual)
	assert.Nil(t, err)
}

func TestPushed_NotTTY(t *testing.T) {
	dummyInput := "a"

	getkbd.UserInputDummy = dummyInput

	defer func() {
		getkbd.UserInputDummy = ""
	}()

	result, err := getkbd.Pushed("q", 0)

	assert.Error(t, err, "non TTY request should return error")
	assert.Empty(t, result, "non TTY request should return empty result and error")
}

func TestPushed_UserInputDummy(t *testing.T) {
	expect := "a"

	getkbd.UserInputDummy = expect
	getkbd.DummyIsTTY = true

	defer func() {
		getkbd.UserInputDummy = ""
		getkbd.DummyIsTTY = false
	}()

	actual, err := getkbd.Pushed("q", 0)
	assert.Equal(t, expect, actual)
	assert.Nil(t, err)
}
