package gochat

type Client interface {
	// StdMessenger is the standard messenger for the client.
	StdMessenger() StdMessenger
}

type StdMessenger interface {
	// SendMessage sends a message to the client.
	SendMessage(message string) (*Response, error)

	// StartThread starts a new thread with the given message. The thread key is the
	// key of the thread to start. The thread key must be unique for each thread.
	// For example, you can use uuid as the thread key.
	//
	// If the thread with the given thread key already exists, the message will be
	// sent to the existing thread.
	StartOrReplyThread(message string, threadKey string) (*Response, error)
}

type Formatter interface {
	Text(text string) Formatter
	Bold(text string) Formatter
	Italic(text string) Formatter
	StrikeThrough(text string) Formatter
	Monospace(text string) Formatter
	Code(text string) Formatter
	List(list []string) Formatter
	NewLine() Formatter
	ToString() string
}
