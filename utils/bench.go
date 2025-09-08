package utils

import (
	"fmt"
	"nso/logging"
	"time"
)

func Bench(f func()) {
	start := time.Now()
	f()
	end := time.Now().Sub(start)
	logging.Logger.Info(fmt.Sprintf("Elapsed time: %d nano seconds\n", end.Nanoseconds()))
}
