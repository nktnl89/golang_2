// Package errors implements functions to demonstrate godoc features
//
// The CallImplicitPanic always returns error with timestamp when panic of dividing by zero was called
//
// CallImplicitPanic() int, err
//
// So, you are allowed to call this function and get error with timestamp of panic
package panic_func

import (
	"fmt"
	"time"
)

type ErrorWithPanic struct {
	text    string
	panicAt time.Time
}

func (e *ErrorWithPanic) Error() string {
	return fmt.Sprintf(e.text)
}

// New create error with timestamp of panic and some description of it
func New(panicTime time.Time) error {
	return &ErrorWithPanic{
		text:    fmt.Sprintf("Panic was detected at: %s", panicTime.Format("2006-01-02 15:04:05")),
		panicAt: panicTime,
	}
}

// CallImplicitPanic always try to return 100 and error that contains time when panic was called
func CallImplicitPanic() (value int, err error) {

	defer func() {
		if v := recover(); v != nil {
			err = New(time.Now())
		}
	}()

	zero := 0
	return 100 / zero, nil
}
