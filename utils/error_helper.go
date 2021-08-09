package utils

import (
	"fmt"
	"runtime"
)

// New error with stack trace support
func NewErrorf(format string, a ...interface{}) error {
	if len(format) > 0 {
		errorMsg := fmt.Sprintf(format, a...)
		_, file, line, _ := runtime.Caller(1)
		return fmt.Errorf("Error [%s:%d] %v", file, line, errorMsg)
	}

	return nil
}
