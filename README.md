# stacktraces

[![GoDoc](https://godoc.org/github.com/songgao/print-stacktraces?status.svg)](http://godoc.org/github.com/songgao/print-stracktraces)

`stacktraces` is a package that makes it easier to bluntly print stacktraces of
all running Go routines on specific signals. Some additional packages are also
included for automatically setting up for `SIGUSR1`, `SIGUSR2`, `SIGTERM`, and
`SIGINT`.

Note that if you use this on `SIGTERM` and `SIGINT`, this would prevent the
process from terminating as in default behavior.

## Example

`/tmp/example/main.go`:

```go
package main

import (
	"time"

	_ "github.com/songgao/print-stracktraces/on/SIGUSR2"
)

func main() {
	time.Sleep(time.Minute)
}
```

```shell
$ go build && ./example
```

Now, send `SIGUSR2` to the process:

```shell
$ killall -SIGUSR2 example
```

The process keeps running but prints following to `stderr`:

```
$ go build && ./example
--- contention:
cycles/second=2900010410
goroutine 7 [running]:
runtime/pprof.writeGoroutineStacks(0x25f1c0, 0xc82002a018, 0x0, 0x0)
	/Users/songgao/go/src/runtime/pprof/pprof.go:516 +0x84
runtime/pprof.writeGoroutine(0x25f1c0, 0xc82002a018, 0x2, 0x0, 0x0)
	/Users/songgao/go/src/runtime/pprof/pprof.go:505 +0x46
runtime/pprof.(*Profile).WriteTo(0x1c1500, 0x25f1c0, 0xc82002a018, 0x2, 0x0, 0x0)
	/Users/songgao/go/src/runtime/pprof/pprof.go:236 +0xd4
github.com/songgao/print-stracktraces.Set.func1(0xc820056060, 0x25f1c0, 0xc82002a018)
	/Users/songgao/gopath/src/github.com/songgao/print-stracktraces/stacktraces.go:71 +0xfa
created by github.com/songgao/print-stracktraces.Set
	/Users/songgao/gopath/src/github.com/songgao/print-stracktraces/stacktraces.go:74 +0xcd

goroutine 1 [sleep]:
time.Sleep(0xdf8475800)
	/Users/songgao/go/src/runtime/time.go:59 +0xf9
main.main()
	/tmp/example/main.go:10 +0x26

goroutine 5 [syscall]:
os/signal.signal_recv(0x25b078)
	/Users/songgao/go/src/runtime/sigqueue.go:116 +0x132
os/signal.loop()
	/Users/songgao/go/src/os/signal/signal_unix.go:22 +0x18
created by os/signal.init.1
	/Users/songgao/go/src/os/signal/signal_unix.go:28 +0x37

goroutine 6 [select, locked to thread]:
runtime.gopark(0x1565c0, 0xc820028728, 0x123260, 0x6, 0x18, 0x2)
	/Users/songgao/go/src/runtime/proc.go:262 +0x163
runtime.selectgoImpl(0xc820028728, 0x0, 0x18)
	/Users/songgao/go/src/runtime/select.go:392 +0xa67
runtime.selectgo(0xc820028728)
	/Users/songgao/go/src/runtime/select.go:215 +0x12
runtime.ensureSigM.func1()
	/Users/songgao/go/src/runtime/signal1_unix.go:279 +0x32c
runtime.goexit()
	/Users/songgao/go/src/runtime/asm_amd64.s:1998 +0x1
heap profile: 1: 1048576 [1: 1048576] @ heap/1048576
1: 1048576 [1: 1048576] @ 0x3f4e5 0x817a5 0x816e6 0x7d064 0x7b53a 0x5a461
#	0x817a5	runtime/pprof.writeGoroutineStacks+0x45			/Users/songgao/go/src/runtime/pprof/pprof.go:514
#	0x816e6	runtime/pprof.writeGoroutine+0x46			/Users/songgao/go/src/runtime/pprof/pprof.go:505
#	0x7d064	runtime/pprof.(*Profile).WriteTo+0xd4			/Users/songgao/go/src/runtime/pprof/pprof.go:236
#	0x7b53a	github.com/songgao/print-stracktraces.Set.func1+0xfa	/Users/songgao/gopath/src/github.com/songgao/print-stracktraces/stacktraces.go:71


# runtime.MemStats
# Alloc = 1161864
# TotalAlloc = 1161864
# Sys = 4458744
# Lookups = 3
# Mallocs = 299
# Frees = 22
# HeapAlloc = 1161864
# HeapSys = 1671168
# HeapIdle = 163840
# HeapInuse = 1507328
# HeapReleased = 0
# HeapObjects = 277
# Stack = 425984 / 425984
# MSpan = 6720 / 16384
# MCache = 4800 / 16384
# BuckHashSys = 1442264
# NextGC = 4194304
# PauseNs = [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
# NumGC = 0
# EnableGC = true
# DebugGC = false
threadcreate profile: total 8
8 @
```
