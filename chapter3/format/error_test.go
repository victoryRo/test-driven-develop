package format_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/victoryRo/test-driven-develop/format"
	"testing"
)

func TestError(t *testing.T) {
	// Arrange
	initialErr := errors.New("error message")
	expr := "2%3"

	// Act
	wrappedErr := format.Error(expr, initialErr)

	// Assert
	assert.Contains(t, wrappedErr.Error(), initialErr.Error())
	assert.Contains(t, wrappedErr.Error(), expr)
}
