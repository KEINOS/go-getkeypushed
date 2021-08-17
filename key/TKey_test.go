package key_test

import (
	"testing"

	"github.com/KEINOS/go-getkeypushed/key"
	"github.com/stretchr/testify/assert"
)

func TestNew_instantiate(t *testing.T) {
	getKey1 := key.New()
	getKey2 := key.New()

	assert.IsType(t, getKey1, getKey2, "created object should be the same type")
	assert.NotSame(t, getKey1, getKey2, "the pointer should not reference the same object")
}

func TestNew_force_default(t *testing.T) {
	getKey := key.New()

	assert.False(t, getKey.ForceDefault, "ForceDefault should be false")

	getKey.ForceDefault = true

	result, err := getKey.Get("key", 0)

	assert.NoError(t, err, "should not return an error")

	expect := "key"
	actual := result
	assert.Equal(t, expect, actual, "should return the default value")
}
