package main

import "syscall"

func main() {
	lws := syscall.NewLazyDLL("User32.dll").NewProc("LockWorkStation")
	lws.Call()
}
