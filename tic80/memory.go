package tic80

import "unsafe"

var Mouse = unsafe.Slice((*byte)(unsafe.Pointer(uintptr(0x0FF84))), 4*1024) // 4KB memory
