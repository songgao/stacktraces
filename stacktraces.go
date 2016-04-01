// Package stacktraces provides a utility function `Set` to make your Go
// program print stack traces of all running Go routines on particularly
// signals.
//
// As an example:
//
//   // file /tmp/example/main.go
//   //
//   package main
//
//   import (
//   	"os"
//   	"syscall"
//   	"time"
//
//   	"github.com/songgao/print-stracktraces"
//   )
//
//   func main() {
//   	stacktraces.Set(os.Stderr, syscall.SIGUSR1)
//   	time.Sleep(time.Minute)
//   }
//
// When the above code is running, send a `SIGUSR1` signal to the process, and it
// will print stacktraces:
//
//   $ killall -SIGUSR1 example
//
// However, in most cases, you'd want this done automatically. This package has
// a few subdirectories (under `on` directory) that have `init()` function
// defined to automatically set up on corresponding signals. For example,
// following is equivelent to the `/tmp/example/main.go` above:
//
//   package main
//
//   import (
//   	"time"
//
//   	_ "github.com/songgao/print-stracktraces/on/SIGUSR1"
//   )
//
//   func main() {
//   	time.Sleep(time.Minute)
//   }
//
package stacktraces

import (
	"io"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
)

func signalSyscalltoOS(syscallSignals []syscall.Signal) (osSignals []os.Signal) {
	for _, s := range syscallSignals {
		osSignals = append(osSignals, os.Signal(s))
	}
	return
}

// Set makes the process catch signals, and write stack traces of all currently
// running Go routines to writer.
func Set(writer io.Writer, signals ...syscall.Signal) {
	c := make(chan os.Signal)
	signal.Notify(c, signalSyscalltoOS(signals)...)
	go func() {
		for range c {
			for _, p := range pprof.Profiles() {
				_ = p.WriteTo(writer, 2)
			}
		}
	}()
}
