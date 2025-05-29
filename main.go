package main

import (
	"fmt"
	"unsafe"

	"github.com/tangzero/tic80.go/tic80"
)

var Mouse = unsafe.Slice((*byte)(unsafe.Pointer(uintptr(0x0FF84))), 4*1024) // 4KB memory

//export TIC
func Tic() {
	// rect(0, 0, 240, 136, 0x000000)   // Clear screen with black
	// rect(10, 10, 220, 116, 0xFFFFFF) // Draw a white rectangle

	tic80.Cls(8)
	tic80.Print("Hello from Go!", 0, 0, 12, true, 1, false)

	x, y := Mouse[0], Mouse[1]
	tic80.Trace(fmt.Sprintf("Mouse position: x=%d, y=%d", x, y), 12)
}

func main() {}
