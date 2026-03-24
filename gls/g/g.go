// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

// Package g exposes goroutine struct g to user space.
package g

import (
	"reflect"
	"unsafe"
)

// getgp returns the pointer to the current runtime.g.
//
//go:nosplit
func getgp() unsafe.Pointer

// getg0 returns an interface to a g struct.
// It uses the current goroutine's g to provide the type information.
//
//go:nosplit
func getg0() interface{} {
	return packEface(getgt(), getgp())
}

// GT returns the type of runtime.g.
//
//go:nosplit
func GT() reflect.Type {
	return getgt()
}

// getgt returns the type of runtime.g.
//
//go:nosplit
func getgt() reflect.Type {
	return typeByString("runtime.g")
}

// G returns current g (the goroutine struct) to user space.
//
//go:nosplit
func G() unsafe.Pointer {
	return getgp()
}

// G0 returns the g0 (main goroutine) or the current goroutine as an interface.
// In newer Go versions, it returns the current goroutine to avoid linker errors
// while still providing the correct type information for reflect.
//
//go:nosplit
func G0() interface{} {
	return getg0()
}
