package calculator_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/victoryRo/test-driven-develop/calculator"
	"testing"
)

func TestGetExpectedLength(t *testing.T) {
	// Arrange
	e := calculator.NewEngine()

	// Act
	opLen := e.GetNumOperands()

	// Assert
	assert.Equal(t, 2, opLen)
}

func TestGetValidOperations(t *testing.T) {
	// Arrange
	e := calculator.NewEngine()

	// Act
	ops := e.GetValidOperators()

	// Assert
	assert.Equal(t, 4, len(ops))
	assert.Contains(t, ops, "+")
	assert.Contains(t, ops, "-")
	assert.Contains(t, ops, "*")
	assert.Contains(t, ops, "/")
}

func TestAdd(t *testing.T) {
	// Arrange
	e := calculator.NewEngine()

	t.Run("positive input", func(t *testing.T) {
		// Arrange
		x, y := 2.5, 3.5
		want := 6.0

		// Act
		result, err := e.Add(x, y)

		// Assert
		require.Nil(t, err)
		require.NotNil(t, result)
		assert.Equal(t, want, *result)
	})
	t.Run("negative input", func(t *testing.T) {
		// Arrange
		x, y := -3.0, -5.7
		want := -8.7

		// Act
		result, err := e.Add(x, y)

		// Assert
		require.Nil(t, err)
		require.NotNil(t, result)
		assert.Equal(t, want, *result)
	})
}

func TestSub(t *testing.T) {
	// Arrange
	e := calculator.NewEngine()

	t.Run("positive input", func(t *testing.T) {
		// Arrange
		x, y := 7.0, 3.0
		want := 4.0

		// Act
		result, err := e.Sub(x, y)

		// Assert
		require.Nil(t, err)
		require.NotNil(t, result)
		assert.Equal(t, want, *result)
	})
	t.Run("negative input", func(t *testing.T) {
		// Arrange
		x, y := -8.0, -3.5
		want := -4.5

		// Act
		result, err := e.Sub(x, y)

		// Assert
		require.Nil(t, err)
		require.NotNil(t, result)
		assert.Equal(t, want, *result)
	})
}

func TestMult(t *testing.T) {
	// Arrange
	e := calculator.NewEngine()

	t.Run("positive input", func(t *testing.T) {
		// Arrange
		x, y := 2.0, 3.0
		want := 6.0

		// Act
		result, err := e.Mult(x, y)

		// Assert
		require.Nil(t, err)
		require.NotNil(t, result)
		assert.Equal(t, want, *result)
	})
	t.Run("negative input", func(t *testing.T) {
		// Arrange
		x, y := 3.0, 3.0
		want := 9.0

		// Act
		result, err := e.Mult(x, y)

		// Assert
		require.Nil(t, err)
		require.NotNil(t, result)
		assert.Equal(t, want, *result)
	})
}

func TestDiv(t *testing.T) {
	// Arrage
	e := calculator.NewEngine()

	t.Run("positive input", func(t *testing.T) {
		// Arrange
		x, y := 20.0, 4.0
		want := 5.0

		// Act
		result, err := e.Div(x, y)

		// Assert
		require.Nil(t, err)
		require.NotNil(t, result)
		assert.Equal(t, want, *result)
	})
	t.Run("negative input", func(t *testing.T) {
		// Arrange
		x, y := -100.0, -5.0
		want := 20.0

		// Act
		result, err := e.Div(x, y)

		// Assert
		require.Nil(t, err)
		require.NotNil(t, result)
		assert.Equal(t, want, *result)
	})
	t.Run("divide by zero", func(t *testing.T) {
		// Arrange
		x, y := 30.0, 0.0
		expectedErr := errors.New("cannot divide by zero")

		// Act
		result, err := e.Div(x, y)

		// Assert
		require.Nil(t, result)
		require.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error())
		assert.Contains(t, err.Error(), expectedErr.Error())
	})
}

func TestProcessOperation(t *testing.T) {
	t.Run("positive operation", func(t *testing.T) {
		// Arrange
		e := calculator.NewEngine()
		op := calculator.Operation{
			Expression: "3 + 3",
			Operator:   "+",
			Operands:   []float64{3, 3},
		}
		expectedResult := "6.0"

		// Act
		result, err := e.ProcessOperation(op)

		// Asseert
		require.Nil(t, err)
		require.NotNil(t, result)
		assert.Contains(t, *result, expectedResult)
		assert.Contains(t, *result, op.Expression)
	})

	t.Run("invalid operation", func(t *testing.T) {
		// Arrange
		e := calculator.NewEngine()
		op := calculator.Operation{
			Expression: "3 % 11",
			Operator:   "%",
			Operands:   []float64{3.0, 11.0},
		}

		// Act
		result, err := e.ProcessOperation(op)

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), op.Expression)
		assert.Contains(t, err.Error(), op.Operator)
	})
}

func BenchmarkAdd(b *testing.B) {
	e := calculator.NewEngine()

	for i := 0; i < b.N; i++ {
		_, _ = e.Add(4, 4)
	}
}

func BenchmarkMult(b *testing.B) {
	e := calculator.NewEngine()

	for i := 0; i < b.N; i++ {
		_, _ = e.Mult(3, 3)
	}
}
