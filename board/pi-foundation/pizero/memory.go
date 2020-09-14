// Raspberry Pi Zero support for tamago/arm
// https://github.com/f-secure-foundry/tamago
//
// Copyright (c) the pizero package authors
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.
//
// +build !linkramsize

package pizero

import (
	// use of go:linkname
	_ "unsafe"
)

//go:linkname ramSize runtime.ramSize
var ramSize uint32 = 0x20000000 - 0x04000000 // 512 MB - 64MB GPU (VideoCore)
