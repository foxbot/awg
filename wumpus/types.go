package wumpus

import (
	"encoding/json"
)

// Event handles a payload from the exchange
type Event struct {
	Op   int             `json:"op"`
	Data json.RawMessage `json:"d"`
	Type string          `json:"t"`
}

// Snowflake is an ID type for Discord
type Snowflake string

// Message handles an incoming message
type Message struct {
	ID        Snowflake `json:"id"`
	ChannelID Snowflake `json:"channel_id"`
	GuildID   Snowflake `json:"guild_id"`
	Content   string    `json:"content"`
}
