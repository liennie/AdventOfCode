package evil

import (
	"fmt"
	"runtime/debug"
)

func Panic(format string, a ...any) {
	panic(fmt.Errorf(format, a...))
}

func Assert(ok bool, format string, a ...any) {
	if !ok {
		panic(fmt.Errorf(format, a...))
	}
}

func Err(err error) {
	if err != nil {
		panic(err)
	}
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
