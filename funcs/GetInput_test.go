package funcs_test

import (
	"errors"
	"testing"

	"github.com/KEINOS/go-getkeypushed/funcs"
	"github.com/mattn/go-tty"
	"github.com/stretchr/testify/assert"
)

func TestGetInput_mock_input(t *testing.T) {
	keyDefault := "q"
	expect := "a"

	funcs.UserInputDummy = expect
	defer func() {
		funcs.UserInputDummy = ""
	}()

	timer := funcs.StartTimer(5)
	defer func() {
		if !timer.Stop() {
			<-timer.C
		}
	}()

	actual, err := funcs.GetInput(timer, keyDefault)

	assert.Equal(t, expect, actual,
		"if funcs.UserInputDummy is not empty then its value should be returned")
	assert.Nil(t, err,
		"if funcs.UserInputDummy is not empty then it should not return an error")
}

func TestGetInput_not_tty(t *testing.T) {
	keyDefault := "q"

	funcs.DummyIsTTY = false
	funcs.UserInputDummy = ""

	OldTTYOpen := funcs.TTYOpen
	defer func() {
		funcs.TTYOpen = OldTTYOpen
	}()

	funcs.TTYOpen = func() (*tty.TTY, error) {
		return nil, errors.New("foo")
	}

	timer := funcs.StartTimer(10)
	defer func() {
		timer.Stop()
	}()

	result, err := funcs.GetInput(timer, keyDefault)

	assert.Error(t, err, "non TTY request should return error")
	assert.Empty(t, result,
		"non TTY request should return empty result and error. Returned msg: %v", result)
}

func TestGetInput_TimeExceed(t *testing.T) {
	keyDefault := "q"

	timer := funcs.StartTimer(5)
	defer func() {
		timer.Stop()
	}()

	result, err := funcs.GetInput(timer, keyDefault)

	assert.Nil(t, err)

	expect := keyDefault
	actual := result
	assert.Equal(t, expect, actual)
}
