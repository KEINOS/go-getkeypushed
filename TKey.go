package getkey

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/mattn/go-isatty"
	"golang.org/x/xerrors"
)

type TKey struct {
	OsExiter   func(int)  // For os.Exit dependency injection
	listener   *TListener // TTY input listener
	fd         uintptr    // For os.Stdout.Fd dependency injection
	ForceDummy bool       // Force dummy mode for test purposes
}

// ExitOnCtrlC closes the tty connection and exits the app
// if interrupt signal is received.
func (k *TKey) ExitOnCtrlC(c chan os.Signal) {
	for sig := range c {
		close(c)

		msgError := "keyboard interrupt"

		if err := k.listener.Close(); err != nil {
			msgError += fmt.Sprintf(" (tty close: %v, signal:%v)", err, sig)
		}

		fmt.Fprintln(os.Stderr, msgError)
		k.OsExiter(1)
	}
}

// Get returns the key pushed in the terminal. If `waitTime` is grather than a 0 (zero)
// then it will wait `waitTime` seconds and returns the `keyDefault` when time exceeds.
func (k *TKey) Get(waitTime int, keyDefault string) (string, error) {
	var (
		key string // the key pushed
		err error  // any error during process
	)

	// Let it work only for terminal usage
	if !k.IsTerminal() {
		return "", xerrors.New("not a terminal")
	}

	// Return the default if forced for testing
	if k.ForceDummy {
		return keyDefault, nil
	}

	// Set channel to receive SIGINT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go k.ExitOnCtrlC(c)

	if waitTime > 0 {
		// Create context with timeout
		ctx, cancel := context.WithTimeout(
			context.Background(),
			time.Duration(waitTime)*time.Second,
		)
		defer cancel()

		key, err = k.GetContext(ctx, keyDefault)
	} else {
		// Create context without timer (works as endless loop)
		ctx := context.Background()

		key, err = k.GetContext(ctx, keyDefault)
	}

	return key, err
}

// GetContext sets channels and returns the result from the Loop method.
func (k *TKey) GetContext(ctxOuter context.Context, keyDefault string) (string, error) {
	// Create context with cancel
	ctxInner, cancel := context.WithCancel(ctxOuter)
	defer cancel()

	// Create tty input listener
	k.listener = NewListener()

	if err := k.listener.Open(); err != nil {
		return "", err
	}
	defer k.listener.Close()

	// Create channel for single key input
	chanInput := make(chan rune)
	defer close(chanInput)

	// Create channel for escaped control codes
	chanEsc := make(chan string)
	defer close(chanEsc)

	// Create channel for error during listening
	errListen := make(chan error)
	defer close(errListen)

	// Run tty input listener in background
	go k.listener.Listen(chanInput, chanEsc, errListen, ctxInner)

	// Loop until it receives any input from the channnels
	for {
		result, err := k.Loop(ctxOuter, ctxInner, keyDefault, chanInput, chanEsc, errListen)
		if !(result == "" && err == nil) {
			return result, err
		}
	}
}

// Loop returns the received results from the channnel.
func (k *TKey) Loop(ctxOuter context.Context, ctxInner context.Context, keyDefault string,
	chanInput chan rune, chanEsc chan string, errListen chan error) (string, error) {
	select {
	// Return on error
	case err := <-errListen:
		if err != nil {
			return "", err
		}
	// Return on timeout
	case <-ctxOuter.Done():
		return keyDefault, nil
	// Return on cancel call
	case <-ctxInner.Done():
		return keyDefault, nil
	// Return on ESC/CTL codes
	case s, ok := <-chanEsc:
		if !ok {
			return "", xerrors.New("failed to get input from tty")
		}

		return s, nil
	// Return on key press
	case r, ok := <-chanInput:
		if !ok {
			return "", xerrors.New("failed to get input from tty")
		}

		return string(r), nil
	}

	// Blank return (do nothing)
	return "", nil
}

// IsTerminal returns true if the app is called from the terminal.
func (k *TKey) IsTerminal() bool {
	if isatty.IsTerminal(k.fd) {
		return true
	}

	if isatty.IsCygwinTerminal(k.fd) {
		return true
	}

	return false
}
