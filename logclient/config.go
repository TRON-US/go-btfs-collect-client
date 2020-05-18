package logclient

import (
	"time"
)

const (
	DebugLevel = iota + 1
	InfoLevel
	WarnLevel
	ErrorLevel
	DevPanicLevel
	PanicLevel
)

const (
	MininumCollectionLogLevel          = 4
	DEFAULT_INPUT_CHANNEL_BUF_CAPACITY = 100
	DEFAULT_NUM_OF_RETRIES             = 3
)

type Configuration struct {
	// label set for the current client.
	Labels string // TODO: <-model.LabelSet

	// The current plan operator tree related parameters

	// URL is the log server http push endpoint
	URL string

	// Destination is the key value of the centralized server into which
	// the current configuration based client is pushing its data.
	// For example, "loki" for a log client to push log events into loki server
	Destination string

	// Transport batch related parameters
	BatchWaitDuration time.Duration
	BatchCapacity     int

	// NetworkOut operator parameters
	NetworkSendTimeout time.Duration
	NetworkSendRetries int
}
