package format_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/victoryRo/test-driven-develop/format"
	"testing"
)

func TestResult(t *testing.T) {
	// Arrange
	expr := "2+3"
	result := 5.55

	// Act
	wrappedResult := format.Result(expr, result)

	// Assert
	assert.Contains(t, wrappedResult, expr)
	assert.Contains(t, wrappedResult, fmt.Sprint(result))
}
