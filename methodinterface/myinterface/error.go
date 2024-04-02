package myinterface

import (
	"fmt"
	"time"
)

type myError struct {
	when time.Time
	what string
}

func (e *myError) Error() string {
	return fmt.Sprintf("at: %v, description: %s", e.when, e.what)
}

func RunError() error {
	return &myError{
		time.Now(),
		"Some error",
	}
}

type errNegativeDiv float64

func (e errNegativeDiv) Error() string {
	return fmt.Sprintf("Cannot devide by %v", float64(e))
}

func Divide(x, y float64) (result float64, err error) {
	if y <= 0 {
		return -1, errNegativeDiv(y)
	}
	return x / y, nil
}
