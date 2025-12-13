package main

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

func main() {
	name := "M Faisal Abdillah\n"

	byte := []byte(name)
	ptr := unsafe.Pointer(&byte[0])

	unix.Syscall(
		syscall.SYS_WRITE, uintptr(1), uintptr(ptr), uintptr(len(name)),
	)
}
