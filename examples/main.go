package main

import (
	"fmt"
	"os"
	"strings"

	getkbd "github.com/KEINOS/go-getkeypushed"
)

func main() {
	var (
		char       string
		err        error
		keyDefault string = "q" // A char to use when wait time exceeds
		timeWait   int    = 5   // Seconds to wait user input
	)

	fmt.Println("Ready. Press any key ... (q = quit)")

	for {
		char, err = getkbd.Pushed(keyDefault, timeWait)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get pressed key. Msg: %v\n", err)
			os.Exit(1)
		}

		if char == "q" {
			fmt.Printf("Key pressed => %#+v\n", char)
			fmt.Println("Quit (q) detected. Exiting ...")
			break
		}

		if strings.TrimSpace(char) == "" {
			fmt.Println("Empty char(white space) detected.")
			break
		}

		fmt.Printf("Key pressed => %#+v\n", char)
	}
}
