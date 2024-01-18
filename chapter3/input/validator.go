package input

import "fmt"

type Validator struct {
	ExpectedLength int
	ValidOperators []string
}

// NewValidator creates a ready to use Validator instance
func NewValidator(expLen int, validOps []string) *Validator {
	return &Validator{
		ExpectedLength: expLen,
		ValidOperators: validOps,
	}
}

func (v *Validator) CheckInput(operator string, operands []float64) error {
	opLen := len(operands)
	if opLen != v.ExpectedLength {
		return fmt.Errorf("unexpected operands length: got %d want %d", opLen, v.ExpectedLength)
	}
	return v.checkOperator(operator)
}

// checkOperator validates the operator is supported
func (v *Validator) checkOperator(operator string) error {
	for _, o := range v.ValidOperators {
		if o == operator {
			return nil
		}
	}
	return fmt.Errorf("invalid operator: %s", operator)
}
