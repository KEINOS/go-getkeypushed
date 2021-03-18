package main

import (
	"fmt"
	"os"
	"strings"

	gkp "github.com/KEINOS/go-getkeypushed"
)

func main() {
	var (
		char string
		err  error
	)

	fmt.Println("Ready. Press any key ...")

	for {
		if char, err = gkp.GetKeyPushed(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get pressed key. Msg: %v\n", err)
			os.Exit(1)
		}

		if strings.TrimSpace(char) == "" {
			fmt.Println("Empty char(white space) detected.")
			break
		}

		fmt.Printf("Key pressed => %#+v\n", char)
	}
}
