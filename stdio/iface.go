package stdio

import "io"

// Flusher interface
type Flusher interface {
	Flush() error
}

// FlushWriter is the interface satisfied by logging destinations.
type FlushWriter interface {
	Flusher
	// Writer the output writer
	io.Writer
}

// FlushCloseWriter is the interface satisfied by logging destinations.
type FlushCloseWriter interface {
	Flusher
	// WriteCloser the output writer
	io.WriteCloser
}

// SyncCloseWriter is the interface satisfied by logging destinations.
// such as os.File
type SyncCloseWriter interface {
	Flusher
	// WriteCloser the output writer
	io.WriteCloser
}
