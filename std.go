package gochat

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	messageReplyOption = "REPLY_MESSAGE_FALLBACK_TO_NEW_THREAD"
)

type StdMessengerImpl struct {
	Client *ClientImpl
}

// StartThread implements StdMessenger.
func (s *StdMessengerImpl) StartOrReplyThread(message string, threadKey string) (*Response, error) {
	resp, err := resty.New().R().
		ForceContentType("application/json").
		SetQueryParam("messageReplyOption", messageReplyOption).
		SetBody(map[string]interface{}{
			"text": message,
			"thread": map[string]string{
				"threadKey": threadKey,
			},
		}).
		SetResult(&Response{}).
		Post(s.Client.WebhookURL)
	if err != nil {
		return nil, err
	}

	resResult, ok := resp.Result().(*Response)
	if !ok {
		return nil, fmt.Errorf("invalid response")
	}

	return resResult, nil
}

// SendMessage implements StdMessenger.
func (s *StdMessengerImpl) SendMessage(message string) (*Response, error) {
	resp, err := resty.New().R().
		ForceContentType("application/json").
		SetBody(map[string]string{
			"text": message,
		}).
		SetResult(&Response{}).
		Post(s.Client.WebhookURL)
	if err != nil {
		return nil, err
	}

	resResult, ok := resp.Result().(*Response)
	if !ok {
		return nil, fmt.Errorf("invalid response")
	}

	return resResult, nil
}

func NewStdMessenger(client *ClientImpl) StdMessenger {
	return &StdMessengerImpl{
		Client: client,
	}
}
