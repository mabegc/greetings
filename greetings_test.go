package greetings_test

import (
	"github.com/mabegc/greetings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	message, err := greetings.Hello("Mabe")

	assert.NotEmpty(t, message)
	assert.NoError(t, err)
}

func TestHello_EmptyName(t *testing.T) {
	message, err := greetings.Hello("")

	assert.Empty(t, message)
	assert.Error(t, err)
}
