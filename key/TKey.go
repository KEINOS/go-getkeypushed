package key

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/mattn/go-isatty"
	"github.com/mattn/go-tty"
	"golang.org/x/xerrors"
)

const CodeESC = 27 // \x1b = \d27= ESC

type TKey struct {
	KeyDefault      string  // The key to be returned if time out occurs
	TTL             int     // The time wait user input in seconds
	DescriptorFile  uintptr // For os.Stdout.Fd dependency injection
	ForceDefault    bool    // If true it will return the default key (for testing)
	DurationDefault int     // The default time wait user input in seconds
}

// ----------------------------------------------------------------------------
//  Functions
// ----------------------------------------------------------------------------

// New returns a new TKey instance pointer with default settings.
func New() *TKey {
	return &TKey{
		ForceDefault:    false,
		KeyDefault:      "",
		TTL:             0,
		DurationDefault: 1 * 60 * 60, // 1 hour
		DescriptorFile:  os.Stdout.Fd(),
	}
}

// ----------------------------------------------------------------------------
//  Public Methods
// ----------------------------------------------------------------------------

// Get returns the key from the user input.
func (k *TKey) Get(keyDefault string, ttl int) (result string, err error) {
	k.KeyDefault = keyDefault
	k.TTL = ttl

	if k.ForceDefault {
		return k.KeyDefault, nil
	}

	// Let it work only for terminal usage
	if !k.isTerminal() {
		return "", xerrors.New("not a terminal")
	}

	ctxCancel, cancel := context.WithCancel(context.Background())
	defer cancel()

	tmpTTY, err := tty.Open()
	if err != nil {
		return "", xerrors.Errorf("failed to open tty: %w", err)
	}
	defer tmpTTY.Close()

	chInput := make(chan string)
	chErr := make(chan error)

	go k.listenTTY(ctxCancel, chInput, chErr, tmpTTY)

	ctxTimer, cancelTimer := context.WithTimeout(context.Background(), k.duration())
	defer cancelTimer()

	chInterrupt := make(chan os.Signal, 1)
	signal.Notify(chInterrupt, os.Interrupt)

	// Loop until user input is received, time out occurs or user cancels
	select {
	case result = <-chInput:
	case err = <-chErr:
	case <-ctxTimer.Done():
		result = k.KeyDefault
	case <-ctxCancel.Done():
		result = k.KeyDefault
	case <-chInterrupt:
		result = ""
		err = xerrors.New("user canceled (ctrl+c, SIGINT detectd)")

		cancel()
	}

	return result, err
}

// ----------------------------------------------------------------------------
//  Private Methods
// ----------------------------------------------------------------------------

// duration returns the time wait user input in seconds.
func (k *TKey) duration() time.Duration {
	if k.TTL <= 0 {
		return time.Duration(k.DurationDefault) * time.Second
	}

	return time.Duration(k.TTL) * time.Second
}

// getBuffer returns the buffered runes from tty appending to buff as a string.
// It will be used to get the whole control code.
func (k *TKey) getBuffer(tmpTTY *tty.TTY, buff string) (string, error) {
	var (
		r   rune
		err error
	)

	// Receive all the buffered chars
	for tmpTTY.Buffered() {
		r, err = tmpTTY.ReadRune()
		if err != nil {
			err = xerrors.Errorf("failed to read from buffer: %v", err)

			break
		}

		buff += string(r)
	}

	return buff, err
}

//  isTerminal returns true if the app is called from the terminal.
//
// For testing purposes, if ForceDefault filesd is true then it will return true.
func (k *TKey) isTerminal() bool {
	if isatty.IsTerminal(k.DescriptorFile) || isatty.IsCygwinTerminal(k.DescriptorFile) || k.ForceDefault {
		return true
	}

	return false
}

// listenTTY is a goroutine that listen to the TTY and read the user input.
func (k *TKey) listenTTY(ctxCancel context.Context, chInput chan string, chErr chan error, tmpTTY *tty.TTY) {
	for {
		r, err := tmpTTY.ReadRune()
		if err != nil {
			chErr <- xerrors.Errorf("failed to read from tty: %v", err)

			break
		}

		if r == 0 {
			continue
		}

		if r != CodeESC {
			chInput <- string(r)

			break
		}

		// If escape code then read the hole control code as a single string
		buff, err := k.getBuffer(tmpTTY, string(r))
		if err != nil {
			chErr <- err
		}

		chInput <- buff

		return
	}

	close(chInput)
	close(chErr)
}
