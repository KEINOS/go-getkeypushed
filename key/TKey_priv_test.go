package key

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTerminal(t *testing.T) {
	getKey := New()

	_, err := getKey.Get("key", 0)

	assert.Error(t, err, "test does not have tty so should return an error")

	getKey.ForceDefault = true

	assert.True(t, getKey.isTerminal(), "it should return true if ForceDefault is true")
}
