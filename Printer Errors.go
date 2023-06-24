package kata

import (
	"fmt"
)

func PrinterError(s string) string {
	errorCount := 0
	for _, char := range s {
		if char < 'a' || char > 'm' {
			errorCount++
		}
	}
	return fmt.Sprintf("%d/%d", errorCount, len(s))
}
