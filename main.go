package main

import (
	"fmt"

	"github.com/tangzero/tic80.go/tic80"
)

//export TIC
func TIC() {
	tic80.Cls(8)
	tic80.Print("Hello from Go!", 10, 10, 12, true, 1, false)

	tic80.Rect(80, 50, 100, 100, 3)
	tic80.Rect(60, 40, 30, 30, 5)

	x, y, left, _, _, _, _ := tic80.Mouse()
	tic80.Print(fmt.Sprintf("Mouse position: x=%d, y=%d", x, y), 10, 20, 12, true, 1, false)
	tic80.Print(fmt.Sprintf("Mouse clicking: %t", left), 10, 30, 12, true, 1, false)

	tic80.Rect(uint32(x), uint32(y), 10, 10, 15)
}

func main() {}
