package indenter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorrectIndentation(t *testing.T) {
	fileString, err := os.ReadFile("./examples/main.c")
	if err != nil {
		return
	}

	res, _ := CorrectIndentation(string(fileString))

	assert.Equal(t, "", res)
}
