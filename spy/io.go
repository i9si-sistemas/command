// Package spy provides a simple writer to capture stdout/stderr for testing or debugging.
package spy

import "io"

// Spy defines an interface for capturing written data.
type Spy interface {
	// Data returns all captured data as a byte slice.
	Data() []byte

	// Write appends data to the internal buffer.
	io.Writer
}

// Writer implements Spy, capturing data into an internal buffer.
type Writer struct {
	buffer []byte
}

// New creates and returns a new Writer instance.
func New() Spy {
	return &Writer{
		buffer: make([]byte, 0),
	}
}

// Data returns the captured output.
func (w *Writer) Data() []byte {
	return w.buffer
}

// Write appends data to the buffer and satisfies the io.Writer interface.
func (w *Writer) Write(p []byte) (n int, err error) {
	w.buffer = append(w.buffer, p...)
	return len(p), nil
}
