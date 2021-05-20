package gkp_test //nolint:testpackage // To override osExit the test needs to be part of main

import (
	"errors"
	"testing"

	gkp "github.com/KEINOS/go-getkeypushed/src"

	"github.com/stretchr/testify/assert"
)

func TestGetKeyPushed_DummyKey(t *testing.T) {
	expect := "a"

	gkp.DummyKey = expect

	defer func() { gkp.DummyKey = "" }()

	actual, err := gkp.GetKeyPushed("q", 5)
	assert.Equal(t, expect, actual)
	assert.Nil(t, err)
}

func TestGetKeyPushed_DummyErr(t *testing.T) {
	expect := "dummy error occurred"

	errDummy := errors.New(expect)

	gkp.DummyErr = errDummy

	defer func() { gkp.DummyErr = nil }()

	msg, err := gkp.GetKeyPushed("q", 5)

	assert.Equal(t, "", msg,
		"When DummyErr global variable is not a nil then GetKeyPushed() should be empty('') message.")

	assert.NotNil(t, err,
		"When DummyErr global variable is not a nil then GetKeyPushed() should return an error.")

	assert.EqualErrorf(t, err, expect, "error message %s", "mal-formatted")
}
