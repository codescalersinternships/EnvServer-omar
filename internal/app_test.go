package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	t.Run("invalid port number", func(t *testing.T) {
		_, err := NewApp(-1)

		assert.Equal(t, err, ErrInvalidPort)
	})

	t.Run("initialize app", func(t *testing.T) {
		_, err := NewApp(8080)

		assert.Nil(t, err)
	})
}
