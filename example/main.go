package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/KEINOS/go-getkeypushed/key"
)

func main() {
	var (
		char       string
		err        error
		keyQuit    = "q"
		keyDefault = "?" // A char to use when wait time exceeds
		timeWait   = 5   // Seconds to wait user input
	)

	keyInput := key.New()

	fmt.Println("Ready. Press any key ... (q = quit)")

	for {
		char, err = keyInput.Get(keyDefault, timeWait)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get pressed key. Msg: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Key pressed => %#+v\n", char)

		if char == keyQuit {
			fmt.Printf("Quit (%s) detected. Exiting ...\n", keyQuit)

			break
		}

		if strings.TrimSpace(char) == "" {
			fmt.Println("Empty char(white space) detected.")

			break
		}
	}
}
