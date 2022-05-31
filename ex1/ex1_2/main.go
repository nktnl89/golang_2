package main

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

func New(panicTime time.Time) error {
	return &ErrorWithPanic{
		text:    fmt.Sprintf("Panic was detected at: %s", panicTime.Format("2006-01-02 15:04:05")),
		panicAt: panicTime,
	}
}

func main() {
	_, err := implicitPanicCall()

	if err != nil {
		fmt.Println(err)
	}

}

func implicitPanicCall() (value int, err error) {
	
	defer func() {
		if v := recover(); v != nil {
			err = New(time.Now())
		}
	}()

	zero := 0
	return 100 / zero, nil
}
