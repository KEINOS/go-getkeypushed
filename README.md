[![golangci-lint](https://github.com/KEINOS/go-getkeypushed/actions/workflows/golangci-lint.yaml/badge.svg)](https://github.com/KEINOS/go-getkeypushed/actions/workflows/golangci-lint.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/go-getkeypushed.svg)](https://pkg.go.dev/github.com/KEINOS/go-getkeypushed#section-documentation "Read generated documentation")

# Go-GetKeyPushed

This package implements `OnKeyPress()` like functionality to the CLI app powered by [`go-tty`](https://github.com/mattn/go-tty/).

It simply returns the key pushed from the console/terminal (TTY) without the `enter` key. If the 1st arg is positive, then it will wait until its time exceeds and returns the 2nd arg as default.

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
keyboard interrupt (tty close: errno 0, signal:interrupt)
exit status 1
```

## Usage

```bash
go get github.com/KEINOS/go-getkeypushed
```

```go
// import getkey "github.com/KEINOS/go-getkeypushed"

key := getkey.New()

fmt.Println("Press any key:")

timeWait := 5 // 5 secs to wait

inputUser, err := key.Get(timeWait, "my default")
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

	getkey "github.com/KEINOS/go-getkeypushed"
)

func main() {
	var (
		char       string
		err        error
		keyDefault = "q" // A char to use when wait time exceeds
		timeWait   = 5   // Seconds to wait user input
	)

	key := getkey.New()

	fmt.Println("Ready. Press any key ... (q = quit)")

	for {
		char, err = key.Get(timeWait, keyDefault)

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

```

## Notes

- Suitable for just getting keypress such as `y`, `n`.
- It is a wrapper of amazing package "[github.com/mattn/go-tty](https://github.com/mattn/go-tty/)" to simplify its usage for my other projects.
  - [github.com/mattn/go-tty](https://github.com/mattn/go-tty/) @ GitHub

## License & Copyright

- [MIT](https://github.com/KEINOS/go-getkeypushed/blob/master/LICENSE)
  - [KEINOS and the contributors](https://github.com/KEINOS/go-getkeypushed/graphs/contributors)
  - [`go-tty`](https://github.com/mattn/go-tty/) & [go-isatty](https://github.com/mattn/go-isatty): Yasuhiro Matsumoto (a.k.a [mattn](https://github.com/mattn/)) @ GitHub
  - [testify](https://github.com/stretchr/testify): [Stretchr](https://github.com/stretchr) @ GitHub
  - See the packages in [go.mod](./go.mod) as well

## TODO

- [ ] Find out how to test the `tty` input
- [x] Implement basic testing
  - [ ] Cover test as much as possible
- [x] Wating time implementaion
  - Wait Nth seconds and return default key if wait time exceed with no user interaction) (Since v2.0)
