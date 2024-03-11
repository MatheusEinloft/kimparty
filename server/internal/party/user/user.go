package user

import (
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	PartyID   string `json:"party_id"`
}

func New(name string, partyID string) *User {
	return &User{
		ID:        uuid.New().String(),
		Name:      name,
		PartyID:   partyID,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}

func (user *User) ToJSON() ([]byte, error) {
	return jsoniter.ConfigFastest.Marshal(user)
}
