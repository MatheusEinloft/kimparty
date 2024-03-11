package message

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/json-iterator/go"
)

type Type int8

const (
	Chat Type = iota
	Event
	Error
	Unknown
)

type Message struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	UserID    string    `json:"user_id"`
	PartyID   string    `json:"party_id"`
	Type      Type      `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type Params struct {
	Content string `json:"content"`
	PartyID string `json:"party_id"`
	UserID  string `json:"user_id"`
	Type    Type   `json:"type"`
}

var messagePool = sync.Pool{
	New: func() interface{} {
		return &Message{}
	},
}

// New gets a new message from the pool and sets the message's fields.
// Don't forget to release the message after using it.
func New(params Params) *Message {
	msg := messagePool.Get().(*Message)
	msg.ID = uuid.New().String()
	msg.Content = params.Content
	msg.UserID = params.UserID
	msg.PartyID = params.PartyID
	msg.Type = params.Type
	msg.CreatedAt = time.Now()
	return msg
}

// It will return the message to the pool.
func (m *Message) ToJSON() ([]byte, error) {
	defer messagePool.Put(m)
	return jsoniter.ConfigFastest.Marshal(m)
}

func MessageFromJSON(bytes []byte) (Message, error) {
	var (
		params Params
		err    error
	)

	err = jsoniter.ConfigFastest.Unmarshal(bytes, &params)
	return *New(params), err
}
