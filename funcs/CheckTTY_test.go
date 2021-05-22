package funcs_test

import (
	"testing"

	"github.com/KEINOS/go-getkeypushed/funcs"
	"github.com/mattn/go-tty"
	"github.com/stretchr/testify/assert"
)

func TestCheckTTY(t *testing.T) {
	chanInput := make(chan string) // Channel for user input
	defer close(chanInput)

	expect := "x"

	funcs.UserInputDummy = expect

	ttyDummy := new(tty.TTY)

	go funcs.CheckTTY(chanInput, ttyDummy)

	actual := <-chanInput

	assert.Equal(t, expect, actual)
}
