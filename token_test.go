package go_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRandomBytes(t *testing.T) {
	b, err := generateRandomBytes(10)
	assert.NoError(t, err)
	assert.Len(t, b, 10)
	assert.IsType(t, []byte{}, b)

	b, err = generateRandomBytes(0)
	assert.NoError(t, err)
	assert.Len(t, b, 0)
}

func TestGenerateRandomString(t *testing.T) {
	b, err := generateRandomString(10)
	assert.NoError(t, err)
	assert.Len(t, b, 10)
	assert.IsType(t, "", b)

	b, err = generateRandomString(0)
	assert.NoError(t, err)
	assert.Len(t, b, 0)
}

func TestGenerateToken(t *testing.T) {
	c := 100000
	a := assert.New(t)

	// Create 100k tokens and ensure they are all different.
	ts := make(map[Token]struct{}, c)

	var token Token
	var err error
	for i := 0; i <= c; i++ {
		token, err = GenerateToken(10)
		a.NoError(err)
		_, ok := ts[token]
		a.False(ok)

		ts[token] = struct{}{}
	}
}
