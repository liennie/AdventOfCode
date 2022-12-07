package evil

import (
	"fmt"
	"runtime/debug"
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
		debug.PrintStack()
	}
}
