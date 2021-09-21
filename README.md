[![golangci-lint](https://github.com/KEINOS/go-getkeypushed/actions/workflows/golangci-lint.yaml/badge.svg)](https://github.com/KEINOS/go-getkeypushed/actions/workflows/golangci-lint.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/go-getkeypushed.svg)](https://pkg.go.dev/github.com/KEINOS/go-getkeypushed/key "Read generated documentation")

# Go-GetKeyPushed

This package implements `OnKeyPress()` like functionality to the CLI app in Golang. (Very much powered by the awesome [`go-tty`](https://github.com/mattn/go-tty/))

It simply returns the key pushed from the console/terminal (TTY) without the `enter` key.

```shellsession
$ # Run sample
$ go run ./example <enter>
Ready. Press any key ... (q = quit)
Key pressed => "a"
Key pressed => "b"
Key pressed => "A"
Key pressed => "B"
Key pressed => "あ"
Key pressed => "い"
Key pressed => "愛"
Key pressed => "\x1b"      // ESC key
Key pressed => "\x1b[A"    // Up arrow key
Key pressed => "\x1b[B"    // Down arrow key
Key pressed => "\x1b[C"    // Right arrow key
Key pressed => "\x1b[D"    // Left arrow key
Key pressed => "\x1b[19~"  // F8 key (if not assigned in the terminal)
Key pressed => "\x1b[20~"  // F9 key (if not assigned in the terminal)
Key pressed => "\x1b[21~"  // F10 key (if not assigned in the terminal)
Key pressed => "q"
Quit (q) detected. Exiting ...
```

```shellsession
$ # Interrupt with Ctrl+c
$ go run ./example <enter>
Ready. Press any key ... (q = quit)
Key pressed => "a"
Key pressed => "b"
Failed to get pressed key. Msg: user cancelled (ctrl+c, SIGINT detectd)
exit status 1
```

## Install

```bash
go get github.com/KEINOS/go-getkeypushed
```

## Usage Sample

```go
// import "github.com/KEINOS/go-getkeypushed/key"

key := key.New() // Instantiate

timeWait := 5              // Timeout user input in 5 secs (0=no timeout)
keyDefault := "my default" // Default value when timeout

fmt.Println("Press any key:")

// The below returns the key pressed from the console/terminal
// or returns the default vaule if it timeouts.
inputUser, err := key.Get(keyDefault, timeWait)
if err != nil {
	fmt.Fprintf(os.Stderr, "Failed to get pressed key. Msg: %v\n", err)
	os.Exit(1)
}

fmt.Println("Input: "+ inputUser)
```

```go
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

```

## Notes

- Develop Online
  - Compatible to [GitHub Codespaces](https://github.com/features/codespaces).
  - Fork this repo and open it via CodeSpaces in your repo.
- Develop Locally
  - No Docker + VSCode user
    - With [GitHub Codespaces](https://marketplace.visualstudio.com/items?itemName=GitHub.codespaces) extension, [you can remotely access to cloud-hosted dev container](https://docs.github.com/ja/codespaces/developing-in-codespaces/using-codespaces-in-visual-studio-code) in GitHub Cosespaces.
  - Docker + VSCode user
    - Compatible to VSCode's [Remote-Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) extension.
    - Fork and clone this repo and "Reopen in Container" via Remote-Containers
    - Or [clone the repo in container volume](vscode://ms-vscode-remote.remote-containers/cloneInVolume?url=https://github.com/KEINOS/go-getkeypushed) for one-time usage.

## License & Copyright

- [MIT](https://github.com/KEINOS/go-getkeypushed/blob/master/LICENSE)
  - [KEINOS and the contributors](https://github.com/KEINOS/go-getkeypushed/graphs/contributors)
  - [`go-tty`](https://github.com/mattn/go-tty/) & [go-isatty](https://github.com/mattn/go-isatty): Yasuhiro Matsumoto (a.k.a [mattn](https://github.com/mattn/)) @ GitHub
  - See the packages in [go.mod](./go.mod) as well

## TODO

- [ ] Find out how to test the `tty` input
- [x] Implement basic testing
  - [ ] Cover test as much as possible
- [x] Wating time implementaion
  - Wait Nth seconds and return default key if wait time exceed with no user interaction)
