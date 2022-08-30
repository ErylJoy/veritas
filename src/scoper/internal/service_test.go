package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	x := Useful()
	assert.Equal(t, "I don't believe it", x)
}
