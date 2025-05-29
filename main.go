package main

import (
	"fmt"

	"github.com/tangzero/tic80.go/tic80"
)

//export TIC
func Tic() {
	tic80.Cls(8)
	tic80.Print("Hello from Go!", 0, 0, 12, true, 1, false)

	tic80.Rect(80, 50, 100, 100, 3)
	tic80.Rect(60, 30, 30, 30, 5)

	x, y := tic80.Mouse[0], tic80.Mouse[1]
	tic80.Trace(fmt.Sprintf("Mouse position: x=%d, y=%d", x, y), 12)
}

func main() {}
