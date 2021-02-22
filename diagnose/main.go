package main

import (
	"fmt"
	"syscall"
	libseccomp "github.com/seccomp/libseccomp-golang"
)

func allowList(syscalls []string) {

	filter, err := libseccomp.NewFilter(libseccomp.ActErrno.SetReturnCode(int16(syscall.EPERM)))
	if err != nil {
		fmt.Printf("Error creating filter: %s\n", err)
	}
	for _, element := range syscalls {
		fmt.Printf("[+] allowlisting: %s\n", element)
		syscallID, err := libseccomp.GetSyscallFromName(element)
		if err != nil {
			panic(err)
		}
		filter.AddRule(syscallID, libseccomp.ActAllow)
	}
	filter.Load()
}

func main() {
	var syscalls = []string{
		"rt_sigaction", "mkdirat", "clone", "mmap", "readlinkat", "futex", "rt_sigprocmask",
		"mprotect", "write", "sigaltstack", "gettid", "read", "open", "close", "fstat", "munmap",
		"brk", "access", "execve", "getrlimit", "arch_prctl", "sched_getaffinity", "set_tid_address", "set_robust_list"}

	allowList(syscalls)

	err := syscall.Mkdir("/tmp/moo", 0755)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("I just created a file\n")
	}
}
