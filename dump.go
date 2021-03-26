package dump

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/davecgh/go-spew/spew"
)

func Dump(a ...interface{}) {
	fDump(os.Stdout, 1, a...)
}

func Fdump(w io.Writer, a ...interface{}) {
	fDump(w, 1, a...)
}

func fDump(w io.Writer, skip int, a ...interface{}) {
	_, file, line, _ := runtime.Caller(skip + 1)
	fmt.Fprintf(w, "------\n%s:%d\n%s", file, line, spew.Sdump(a...))
}
