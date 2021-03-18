# Go-GetKeyPushed

**It simply gets one character of the key pushed** from the console (TTY) without entering the `enter` key. Some what like `OnKeyPress()` functionality.

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
Empty char(white space) detected. Exiting ...
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
        char string
        err  error
    )

    fmt.Println("Ready. Press any key ...")

    for {
        if char, err = gkp.GetKeyPushed(); err != nil {
            fmt.Fprintf(os.Stderr, "Failed to get key. Msg: %v\n", err)
            os.Exit(1)
        }

        if strings.TrimSpace(char) == "" {
            fmt.Println("Empty char(white space) detected. Exiting ...")
            break
        }

        fmt.Printf("Key pressed => %#+v\n", char)
    }
}
```

## Tips For Testing

Since testing the `tty` input is difficult, you can mock the `gkp.GetKeyPushed()` behavior by changing the global variables such as `gkp.DummyMsg` and `gkp.DummyErr`.

If you set any string to `gkp.DummyMsg` then `gkp.GetKeyPushed()` will return that string instead of `tty` input.

```go
func TestGetKeyPushed_DummyMsg(t *testing.T) {
    expect := "a"

    gkp.DummyMsg = expect

    defer func() { gkp.DummyMsg = "" }()

    actual, err := gkp.GetKeyPushed()
    assert.Equal(t, expect, actual)
    assert.Nil(t, err)
}
```

If you set an `error` to `gkp.DummyErr` then `gkp.GetKeyPushed()` will return that error.

```go
func TestGetKeyPushed_DummyErr(t *testing.T) {
    expect := "dummy error occured"

    errDummy := errors.New(expect)

    gkp.DummyErr = errDummy

    defer func() { gkp.DummyErr = nil }()

    msg, err := gkp.GetKeyPushed()

    assert.EqualErrorf(t, err, expect, "error message %s", "mal-formatted")
}
```

## Notes

- Suitable for just getting keypress such as `y`, `n`. But not suitable for capturing keys like `F1` nor `SIGINT` (`Ctrl+C`).
- It is a wrapper of amazing package "[github.com/mattn/go-tty](https://github.com/mattn/go-tty/)" to simplify its usage for my other projects.
  - [github.com/mattn/go-tty](https://github.com/mattn/go-tty/) @ GitHub

## License & Copyright

- [MIT](https://github.com/KEINOS/go-getkeypushed/blob/master/LICENSE)
- Author of [`go-tty`](https://github.com/mattn/go-tty/): Yasuhiro Matsumoto (a.k.a [mattn](https://github.com/mattn/))
- [KEINOS and the contributors](https://github.com/KEINOS/go-getkeypushed/graphs/contributors)

## TODO

- [ ] Find out how to test the `tty` input
- [ ] Implement testing
