//go:build linux

package corplink

import (
	"fmt"
	"os"
	"syscall"
)

func HandleParentDeathSignal() {
	_, _, e := syscall.Syscall(syscall.SYS_PRCTL, syscall.PR_SET_PDEATHSIG, uintptr(syscall.SIGTERM), 0)
	if e != 0 {
		_, _ = fmt.Fprintf(os.Stderr, "prctl PR_SET_PDEATHSIG fail: %s", e.Error())
	}
}
