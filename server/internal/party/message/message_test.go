package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessage_ParamsFromJSON(t *testing.T) {
	json := `{"party_id":"123","user_id":"123","content":"test message","type":0}`

	msg, err := MessageFromJSON([]byte(json))

	assert.Nil(t, err)
	assert.Equal(t, "123", msg.PartyID)
	assert.Equal(t, "123", msg.UserID)
	assert.Equal(t, "test message", msg.Content)
	assert.Equal(t, Type(0), msg.Type)
}
