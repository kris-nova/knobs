package stream

type StreamerOptions struct {
	StreamListenAddr string
}

type Streamer interface {

	// Stream is the main synchronous method
	// for streaming based on a protocol defined
	// in an underlying implementation.
	Stream() error
}
