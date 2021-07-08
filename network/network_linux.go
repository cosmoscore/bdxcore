// +build linux
package network

import (
	"fmt"
	"runtime"
	"syscall"
)

func SetLimit() {
	if runtime.GOOS == "linux" {
		var rLimit syscall.Rlimit
		if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
			panic(err)
		}
		rLimit.Cur = rLimit.Max
		if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
			panic(err)
		}

		fmt.Printf("set cur limit: %d\n", rLimit.Cur)
	}
}
