package gkp

import (
	"time"

	"github.com/mattn/go-tty"
)

func getInput(timer *time.Timer, keyDefault string) (string, error) {
	var (
		t      *tty.TTY
		result string
		err    error
	)

	if t, err = OpenTTY(); err != nil {
		return "", err
	}
	defer t.Close()

	chanInput := make(chan string) // Channel for user input
	defer close(chanInput)

	for {
		go checkTime(chanInput, timer, keyDefault)
		go checkTTY(chanInput, t)

		if inputViaChan := <-chanInput; inputViaChan != "" {
			result = inputViaChan

			break
		}

		time.Sleep(1 * time.Second)
	}

	return result, nil
}
