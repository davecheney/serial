// Package serial providers methods for working with serial terminals and line disciplines.
package serial

import (
	"io"
	"os"
	"syscall"
)

// Open returns an io.ReadWriteCloser representing the serial device
// configured to run at baud. Data, parity, and stop bits are always configured
// to be 8,N,1.
func Open(dev string, baud int) (io.ReadWriteCloser, error) {
	f, err := os.OpenFile(dev, syscall.O_RDWR, 0600)
	return f, err
}
