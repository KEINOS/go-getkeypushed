package funcs

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-tty"
)

func CheckTTY(returnChan chan<- string, t *tty.TTY) {
	if UserInputDummy != "" {
		returnChan <- strings.TrimSpace(UserInputDummy)

		return
	}

	r, err := t.ReadRune()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open TTY in checkTTY().\nErrMsg: %v\n", err)

		OsExit(1)
	}

	returnChan <- strings.TrimSpace(string(r))
}
