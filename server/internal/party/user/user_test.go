package user

import (
    "testing"
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestUser_ToJSON(t *testing.T) {
	user := New("test", "test")

	json, err := user.ToJSON()

	assert.Nil(t, err)
	assert.NotNil(t, json)

	parsedJSON := &User{}

	err = jsoniter.ConfigFastest.Unmarshal(json, parsedJSON)

	assert.Nil(t, err)
	assert.Equal(t, user.ID, parsedJSON.ID)
	assert.Equal(t, user.Name, parsedJSON.Name)
	assert.Equal(t, user.PartyID, parsedJSON.PartyID)
	assert.Equal(t, user.CreatedAt, parsedJSON.CreatedAt)
}
