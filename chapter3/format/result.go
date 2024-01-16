package format

import "fmt"

// Result of operation
func Result(expr string, result float64) string {
	return fmt.Sprintf("CALCULATION SUCCESS: %s = %.2f", expr, result)
}
