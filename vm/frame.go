package vm

import (
	"go-compiler/code"
	"go-compiler/object"
)

type Frame struct {
	cl          *object.Closure
	ip          int
	basePointer int
}

func NewFrame(fn *object.Closure, basePointer int) *Frame {
	return &Frame{cl: fn, ip: -1, basePointer: basePointer}
}

func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}
