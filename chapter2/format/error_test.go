package format_test

import (
	"errors"
	"github.com/victoryRo/test-driven-develop/format"
	"strings"
	"testing"
)

func TestError(t *testing.T) {
	// Arrange
	initialErr := errors.New("error msg")
	expr := "2%3"

	// Act
	wrappedErr := format.Error(expr, initialErr)

	// Assert
	if !strings.Contains(wrappedErr.Error(), initialErr.Error()) {
		t.Errorf("error does not contain: got: %s, want: %s", wrappedErr.Error(), initialErr.Error())
	}
	if !strings.Contains(wrappedErr.Error(), expr) {
		t.Errorf("error does not contain: got: %s, want: %s", wrappedErr.Error(), expr)
	}
}
