package main

import (
	"github.com/tangzero/tic80.go/tic80"
)

//export TIC
func TIC() {
	tic80.Cls(8)
	tic80.Print("Hello from Go!", 10, 10, 12, true, 1, false)

	tic80.Rect(80, 50, 100, 100, 3)
	tic80.Rect(60, 40, 30, 30, 5)

	x, y, left, _, _, _, _ := tic80.Mouse()
	tic80.Printf("Mouse position: x=%d y=%d", 10, 20, 12, true, 1, false, x, y)
	tic80.Printf("Mouse clicking: %t", 10, 30, 12, true, 1, false, left)

	color := uint32(15)
	if left {
		color = 10 // Change color if left mouse button is pressed
	}
	tic80.Rect(uint32(x)-10, uint32(y)-10, 20, 20, color)
}

// intentional empty main function to satisfy Go's requirement
func main() {}
