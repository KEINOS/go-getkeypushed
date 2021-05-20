[![golangci-lint](https://github.com/KEINOS/go-getkeypushed/actions/workflows/golangci-lint.yaml/badge.svg)](https://github.com/KEINOS/go-getkeypushed/actions/workflows/golangci-lint.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/go-getkeypushed.svg)](https://pkg.go.dev/github.com/KEINOS/go-getkeypushed#section-documentation "Read generated documentation")

# Go-GetKeyPushed

It simply gets a character of the key pushed from the console/terminal (TTY).

No need to enter the `enter` key. Some what like `OnKeyPress()` functionality. It is very much powered by [`go-tty`](https://github.com/mattn/go-tty/).

```shellsession
$ go run ./examples/main.go <enter>
Ready. Press any key ...
Key pressed => "a"
Key pressed => "b"
Key pressed => "A"
Key pressed => "B"
Key pressed => "あ"
Key pressed => "い"
Key pressed => "\x1b" // Up arrow
Key pressed => "\x1b" // Down arrow
Key pressed => "q"
Quit (q) key detected. Exiting ...
$
```

## Usage

```bash
go get github.com/KEINOS/go-getkeypushed
```

```go
package main

import (
    "fmt"
    "os"
    "strings"

    gkp "github.com/KEINOS/go-getkeypushed"
)

func main() {
    var (
        timeWait int = 10       // Wait 10 seconds ( 0 = wait for ever)
        keyDefault string = "q" // Key to use when wait time exceeds
    )

    fmt.Println("Ready. Press any key ...")

    for {
        char, err = gkp.GetKeyPushed(keyDefault, timeWait)

        if err != nil {
            fmt.Fprintf(os.Stderr, "Failed to get key. ErrMsg: %v\n", err)
            os.Exit(1)
        }

        if char == "q" {
            fmt.Printf("Key pressed => %#+v\n", char)
            fmt.Println("Quit (q) key detected. Exiting ...")
            break
        }

        fmt.Printf("Key pressed => %#+v\n", char)
    }
}
```

## Notes

- Suitable for just getting keypress such as `y`, `n`. But not suitable for capturing keys like `F1` nor `SIGINT` (`Ctrl+C`).
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
- [x] Wating time implementaion
  - Wait Nth seconds and return default key if wait time exceed with no user interaction) (Since v2.0)
