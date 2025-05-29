package tic80

import (
	"runtime"
	"unsafe"
)

//go:wasmimport env print
func print(text uint32, x uint32, y uint32, color uint32, fixed bool, scale uint32, smallfont bool) uint32

func Print(text string, x uint32, y uint32, color uint32, fixed bool, scale uint32, smallfont bool) uint32 {
	ptr := stringToPtr(text)
	ret := print(ptr, x, y, color, fixed, scale, smallfont)
	runtime.KeepAlive(text)
	return ret
}

//go:wasmimport env trace
func trace(text, color uint32)

func Trace(text string, color uint32) {
	ptr := stringToPtr(text)
	trace(ptr, color)
	runtime.KeepAlive(text)
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
