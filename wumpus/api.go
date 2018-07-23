package wumpus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Discord wraps the Discord REST API (blaze)
//
// Note: This wrapper does NOT provide rate-limiting, you are expected to be using
// a proxy ala. blaze which provides consolidated ratelimiting
type Discord struct {
	baseURL string
	client  *http.Client
}

// NewDiscord creates a new Discord wrapper at the given URL
func NewDiscord(baseURL string) *Discord {
	var client http.Client
	return &Discord{
		baseURL: baseURL,
		client:  &client,
	}
}

func (discord *Discord) makeRequest(method, endpoint string, body interface{}) (*http.Request, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(b)
	req, err := http.NewRequest(method, endpoint, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "DiscordBot (github.com/foxbot/awg/wumpus, r0)")
	return req, nil
}

// CreateMessageArgs contains additional fields to create a message (TODO)
type CreateMessageArgs struct {
	Content string      `json:"content",omitempty`
	Embed   interface{} `json:"embed",omitempty`
}

func (d *Discord) CreateMessage(channelID string, args *CreateMessageArgs) (*Message, error) {
	url := fmt.Sprintf("channels/%s/messages", channelID)
	req, err := d.makeRequest("POST", url, args)
	if err != nil {
		return nil, err
	}

	r, err := d.client.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}

	var m Message
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
