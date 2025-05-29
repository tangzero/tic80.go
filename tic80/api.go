package tic80

import (
	"fmt"
	"runtime"
	"unsafe"
)

//go:wasmimport env mouse
func mouse(data uint32)

type mouseData struct {
	x       int16
	y       int16
	scrollx int8
	scrolly int8
	left    bool
	middle  bool
	right   bool
}

func Mouse() (x, y int16, left, middle, right bool, scrollX, scrollY int8) {
	var data mouseData
	mouse(uint32(uintptr((unsafe.Pointer(&data)))))
	return data.x, data.y, data.left, data.middle, data.right, data.scrollx, data.scrolly
}

//go:wasmimport env print
func print(text uint32, x uint32, y uint32, color uint32, fixed bool, scale uint32, smallfont bool) uint32

func Print(text string, x uint32, y uint32, color uint32, fixed bool, scale uint32, smallfont bool) uint32 {
	ptr := stringToPtr(text)
	ret := print(ptr, x, y, color, fixed, scale, smallfont)
	runtime.KeepAlive(text)
	return ret
}

func Printf(format string, x uint32, y uint32, color uint32, fixed bool, scale uint32, smallfont bool, args ...any) uint32 {
	return Print(fmt.Sprintf(format, args...), x, y, color, fixed, scale, smallfont)
}

//go:wasmimport env trace
func trace(text, color uint32)

func Trace(text string, color uint32) {
	ptr := stringToPtr(text)
	trace(ptr, color)
	runtime.KeepAlive(text)
}

func Tracef(format string, color uint32, args ...any) {
	Trace(fmt.Sprintf(format, args...), color)
}

//go:wasmimport env rect
func Rect(x, y, w, h, color uint32)

//go:wasmimport env time
func Time() float32

//go:wasmimport env cls
func Cls(color uint32)

//go:wasmimport env exit
func Exit()

func stringToPtr(s string) uint32 {
	buf := []byte(s + "\x00")
	ptr := &buf[0]
	return uint32(uintptr(unsafe.Pointer(ptr)))
}
