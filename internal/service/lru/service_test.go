package lru

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testSize = 3
)

func TestService(t *testing.T) {
	s := NewService(testSize)
	firstKey := "k1"
	for i := 1; i <= testSize; i++ {
		ok := s.Add("k"+strconv.Itoa(i), "v"+strconv.Itoa(i))
		assert.True(t, ok)
	}

	ok := s.Add("k", "v")
	assert.True(t, ok)

	value, ok := s.Get(firstKey)
	assert.False(t, ok)
	assert.Equal(t, "", value)

	value, ok = s.Get("k")
	assert.True(t, ok)
	assert.Equal(t, "v", value)

	ok = s.Remove("k1")
	assert.False(t, ok)

	ok = s.Remove("k")
	assert.True(t, ok)
	ok = s.Remove("k")
	assert.False(t, ok)
}
