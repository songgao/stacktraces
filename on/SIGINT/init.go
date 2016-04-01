package SIGINT

import (
	"os"
	"syscall"

	"github.com/songgao/print-stacktraces"
)

func init() {
	stacktraces.Set(os.Stderr, syscall.SIGINT)
}
