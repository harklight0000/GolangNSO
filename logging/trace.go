package logging

import (
	"fmt"
	"runtime"
)

type Caller struct {
	Line int
	File string
}

func NewCaller() *Caller {
	_, file, line, _ := runtime.Caller(2)
	return &Caller{
		Line: line,
		File: convertToShortFile(file),
	}
}

func (c *Caller) String() string {
	return fmt.Sprintf("%s:%d", c.File, c.Line)
}

func convertToShortFile(file string) string {
	runes := []rune(file)
	count := 0
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == '/' {
			if count == 1 {
				return string(runes[i+1:])
			} else {
				count++
			}
		}
	}
	return file
}
