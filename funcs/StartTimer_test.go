package funcs_test

import (
	"testing"

	"github.com/KEINOS/go-getkeypushed/funcs"
	"github.com/stretchr/testify/assert"
)

func TestStartTimer(t *testing.T) {
	timer1 := funcs.StartTimer(100)
	timer2 := funcs.StartTimer(100)

	defer func() {
		if !timer1.Stop() {
			<-timer1.C
		}

		if !timer2.Stop() {
			<-timer2.C
		}
	}()

	assert.IsType(t, timer1, timer2, "returned instances should be the same type")
	assert.NotSame(t, timer1, timer2, "the returned object pointer should not point the same object")
}
