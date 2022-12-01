package util

import (
	"fmt"
)

func Panic(format string, a ...interface{}) {
	panic(fmt.Errorf(format, a...))
}

func Recover(f func(err error)) {
	if e := recover(); e != nil {
		switch e := e.(type) {
		case error:
			f(e)
		default:
			f(fmt.Errorf("%v", e))
		}
	}
}
