package getkey

import (
	"os"
)

func New() *TKey {
	return &TKey{
		fd: os.Stdout.Fd(),
	}
}
