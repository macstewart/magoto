package logger

import (
	"fmt"
	"time"
)

func Timer(caller string, args ...any) func() {
	caller = fmt.Sprintf(caller, args...)
	start := time.Now()
	return func() {
		Trace("(TIMER) %s took %s", caller, time.Since(start))
	}
}
