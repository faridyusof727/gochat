# Go Chat

A simple Go library for interacting with Google Chat API.

## Installation

```bash
go get github.com/faridyusof727/gochat@v0.1.3
```

## Features

- Send messages to Google Chat
- Start or reply to threads
- Simple and easy-to-use interface
- Custom format of your message (v0.1.3)
  
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

### Send Simple Message with formatting

```go
// Create a formatter
f := gochat.NewFormatter()
    f.Bold("Main Title").
        NewLine().
        Italic("Username: johndoe").
        NewLine().
        NewLine().
        List([]string{"Name: John Doe", "Age: 20", "Category: 2A"}).
        NewLine().
        Code("package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, world!\")")

// Send message with formatting
response, err := client.StdMessenger().SendMessage(f.ToString())
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
