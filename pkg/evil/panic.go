package evil

import (
	"errors"
	"fmt"
	"runtime/debug"
)

func Panic(format string, a ...interface{}) {
	panic(fmt.Errorf(format, a...))
}

func Assert(ok bool, a ...any) {
	if !ok {
		panic(errors.New(fmt.Sprint(a...)))
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
