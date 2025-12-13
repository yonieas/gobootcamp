package main

import (
	"syscall"
	"unsafe"
)

func main () {
	// string yang bakal di cetak
	name := "Herbayu"

	// ini di split kaya di js (cuman ini berdasarkan ASCII)
	value := []byte(name)

	// stdout itu 1 (standar output)
	fd := 1

	// pointer ke buffer value nya
	buf := unsafe.Pointer(&value[0])

	// panjang dari value
	lengthName := len(value)

	syscall.Syscall(syscall.SYS_WRITE, uintptr(fd), uintptr(buf), uintptr(lengthName))
}