package cmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConMap_Get(t *testing.T) {
	cm := New[int]()

	cm.Set("test", 1)

	value, ok := cm.Get("test")

	assert.True(t, ok)
	assert.Equal(t, 1, value)
}

func TestConMap_Set(t *testing.T) {
	cm := New[int]()

	cm.Set("test", 1)

	value, ok := cm.Get("test")

	assert.True(t, ok)
	assert.Equal(t, 1, value)
}

func TestConMap_Remove(t *testing.T) {
	cm := New[int]()

	cm.Set("test", 1)
	cm.Remove("test")

	_, ok := cm.Get("test")

	assert.False(t, ok)
}

func TestConMap_Count(t *testing.T) {
	cm := New[int]()

	cm.Set("test", 1)
	count := cm.Count()

	assert.Equal(t, 1, count)
}

func TestConMap_IterWithKey(t *testing.T) {
	cm := New[int]()

	cm.Set("test", 1)

	for tuple := range cm.IterWithKey() {
		assert.Equal(t, "test", tuple.Key)
		assert.Equal(t, 1, tuple.Val)
	}
}

func TestConMap_Iter(t *testing.T) {
	cm := New[int]()

	cm.Set("test", 1)

	for value := range cm.Iter() {
		assert.Equal(t, 1, value)
	}
}

func TestConMap_Values(t *testing.T) {
	cm := New[int]()

	cm.Set("test", 1)

	values := cm.Values()

	assert.Equal(t, 1, values[0])
}
