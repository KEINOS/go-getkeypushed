package funcs

import (
	"time"

	"github.com/mattn/go-tty"
)

func GetInput(timer *time.Timer, keyDefault string) (string, error) {
	var (
		t      *tty.TTY
		result string
		err    error
	)

	if UserInputDummy != "" {
		return UserInputDummy, nil
	}

	if t, err = OpenTTY(); err != nil {
		return "", err
	}
	defer t.Close()

	chanInput := make(chan string) // Channel for user input
	defer close(chanInput)

	for {
		go CheckTime(chanInput, timer, keyDefault)
		go CheckTTY(chanInput, t)

		if inputViaChan := <-chanInput; inputViaChan != "" {
			result = inputViaChan

			break
		}
	}

	return result, nil
}
