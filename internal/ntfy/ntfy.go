package ntfy

import (
	"fmt"
	"net/http"
	"strings"
)

type Client struct {
	Url        string
	AuthHeader string
}

func NewClient(baseUrl string, topic string, token string) *Client {
	client := &Client{
		Url:        fmt.Sprintf("%s/%s", baseUrl, topic),
		AuthHeader: fmt.Sprintf("Bearer %s", token),
	}

	return client
}

func (c Client) SendOccupationNotification(occupied bool) {
	var message string
	if occupied {
		message = "Bathroom is occupied 🚽"
	} else {
		message = "Bathroom is not occupied 😃"
	}

	req, _ := http.NewRequest("POST", c.Url, strings.NewReader(message))
	req.Header.Set("Authorization", c.AuthHeader)
	http.DefaultClient.Do(req)
}
