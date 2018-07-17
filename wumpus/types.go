package wumpus

import (
	"encoding/json"
)

// Event handles a payload from the sharder
type Event struct {
	Op   int             `json:"op"`
	Data json.RawMessage `json:"d"`
	Type string          `json:"string"`
}
