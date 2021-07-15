package getkey_test

import (
	"os"
	"os/signal"
	"testing"

	getkey "github.com/KEINOS/go-getkeypushed"
	"github.com/stretchr/testify/assert"
)

func TestExitOnCtrlC(t *testing.T) {
	keyTmp := getkey.New()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var status int

	keyTmp.OsExiter = func(exitCode int) {
		status = exitCode
	}

	assert.NotPanics(t, func() {
		go keyTmp.ExitOnCtrlC(c)

		c <- os.Interrupt
	})
	assert.Equal(t, 0, status)
}
