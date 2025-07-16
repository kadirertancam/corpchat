package chat

import "time"

type Message struct {
	ID        int64     `json:"id"`
	FromUID   int64     `json:"from_uid"`
	ToUID     int64     `json:"to_uid,omitempty"`
	ChannelID *int      `json:"channel_id,omitempty"`
	Body      string    `json:"body"`
	FileURL   *string   `json:"file_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
