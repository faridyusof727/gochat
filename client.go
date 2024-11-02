package gochat

import (
	"fmt"
	"net/url"
)

type Option func(*ClientImpl) error

type ClientImpl struct {
	WebhookURL string
}

// StdMessenger implements Client.
func (c *ClientImpl) StdMessenger() StdMessenger {
	return NewStdMessenger(c)
}

func New(opts ...Option) (Client, error) {
	client := &ClientImpl{
		WebhookURL: "",
	}

	for _, opt := range opts {
		err := opt(client)
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}

func WithWebhookURL(hookUrl string) func(*ClientImpl) error {
	return func(c *ClientImpl) error {
		_, err := url.ParseRequestURI(hookUrl)
		if err != nil {
			return err
		}

		c.WebhookURL = hookUrl
		return nil
	}
}

func WithConfig(space string, key string, token string) func(*ClientImpl) error {
	return func(c *ClientImpl) error {
		c.WebhookURL = fmt.Sprintf("https://chat.googleapis.com/v1/spaces/%s/messages?key=%s&token=%s", space, key, token)
		return nil
	}
}
