package main

import (
	"fmt"
	"os"

	gkp "github.com/KEINOS/go-getkeypushed/src"
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
		if char, err = gkp.GetKeyPushed(keyDefault, timeWait); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get pressed key. Msg: %v\n", err)
			os.Exit(1)
		}

		if char == "" {
			fmt.Println("Empty char(white space) detected.")
			break
		}

		if char == "q" {
			fmt.Printf("Key pressed => %#+v\n", char)
			fmt.Println("Quit (q) detected. Exiting ...")
			break
		}

		fmt.Printf("Key pressed => %#+v\n", char)
	}
}
