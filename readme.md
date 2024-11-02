# Go Chat

A simple Go library for interacting

## Installation

```bash
go get github.com/faridyusof727/go-chat
```

## Features

- Send messages to Google Chat
- Start or reply to threads
- Simple and easy-to-use interface

## Usage

### Initialize Client

```go
client, err := gochat.New(
    gochat.WithConfig(
        "YOUR_SPACE_ID",
        "YOUR_API_KEY",
        "YOUR_TOKEN",
    ),
)
if err != nil {
    log.Fatal(err)
}

// Or...

client, err := gochat.New(
    gochat.WithWebhookURL("<Your webhook url>"),
)
if err != nil {
    log.Fatal(err)
}
```

### Send Simple Message

```go
response, err := client.StdMessenger().SendMessage("Hello, world!")
if err != nil {
    log.Fatal(err)
}
```

### Send Message in Thread

```go
// Start or reply to a thread
response, err := client.StdMessenger().StartOrReplyThread(
    "Hello, world!",
    "unique-thread-key",
)
if err != nil {
    log.Fatal(err)
}

// Send another message in the same thread
response, err = client.StdMessenger().StartOrReplyThread(
    "This is a reply",
    "unique-thread-key",
)
if err != nil {
    log.Fatal(err)
}
```

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
