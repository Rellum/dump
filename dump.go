package dump

import (
	"io"
	"os"
	"runtime"

	"github.com/davecgh/go-spew/spew"
)

func Dump(a ...interface{}) {
	Fdump(os.Stdout, a...)
}

func Fdump(w io.Writer, a ...interface{}) {
	caller, file, line, ok := runtime.Caller(1)
	spew.Fdump(w, append([]interface{}{caller, file, line, ok}, a...)...)
}
