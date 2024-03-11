package party

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestParty_ToJSON(t *testing.T) {
	pt := New("http://example.com", 10)

	json, err := pt.ToJSON()

	assert.Nil(t, err)
	assert.NotNil(t, json)

	parsedJSON := &Party{}

	err = jsoniter.ConfigFastest.Unmarshal(json, parsedJSON)

	assert.Nil(t, err)
	assert.Equal(t, pt.ID, parsedJSON.ID)
	assert.Equal(t, pt.URL, parsedJSON.URL)
	assert.Equal(t, pt.CreatedAt, parsedJSON.CreatedAt)
}
