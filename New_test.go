package getkey_test

import (
	"testing"

	getkey "github.com/KEINOS/go-getkeypushed"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	obj1 := getkey.New()
	obj2 := getkey.New()

	assert.IsType(t, obj1, obj2, "new objects should be the same type")
	assert.NotSame(t, obj1, obj2, "the pointers should not reference the same object")
}
