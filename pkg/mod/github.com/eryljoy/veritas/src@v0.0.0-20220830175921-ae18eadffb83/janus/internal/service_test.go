package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHahaha(t *testing.T) {
	assert.Equal(t, "perfect", ReadFile())
}
