package getkey

import (
	"context"

	"github.com/mattn/go-tty"
)

const CodeESC = 27 // \x1b = \d27= ESC

type TListener struct {
	tty *tty.TTY
}

// The newListener returns an instantiated TListener object.
func NewListener() *TListener {
	return &TListener{
		tty: new(tty.TTY),
	}
}

// Close closes the connection to tty.
func (l *TListener) Close() (err error) {
	return l.tty.Close()
}

// Open opens connection to tty.
func (l *TListener) Open() (err error) {
	l.tty, err = tty.Open()

	return err
}

// Listen begins to read input from the tty connection and returns the result to the channels in the args.
func (l *TListener) Listen(inputKey chan rune, escCode chan string, errListen chan error, ctx context.Context) {
	// Begin read tty
	r, err := l.Read()
	if err != nil {
		errListen <- err

		return
	}

	// Escape control codes
	if r == CodeESC {
		buff := string(r)

		// Receive all the buffered chars
		for l.tty.Buffered() {
			r, err = l.Read()
			if err != nil {
				errListen <- err

				return
			}

			buff += string(r)
		}

		escCode <- buff

		return
	}

	// Other chars
	inputKey <- r
}

// Read begins to read a single char(rune) from the tty connection.
func (l *TListener) Read() (rune, error) {
	return l.tty.ReadRune()
}
