package dump

import (
	"runtime"

	"github.com/davecgh/go-spew/spew"
)

func Dump(a ...interface{}) {
	caller, file, line, ok := runtime.Caller(1)
	spew.Dump(append([]interface{}{caller, file, line, ok}, a...)...)
}
