package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToekn(t *testing.T) {
	token, err := GenerateToken("1")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}