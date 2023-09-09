package main

import (
	"testing"
)

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
	Thursday
)

const (
	Executable = 1 << iota
	Writable
	Readable
)

func TestFibonacci(t *testing.T) {
	var a = 1
	var b = 1
	t.Log(" ", a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a + b
		a = b
		b = tmp
	}
}

func TestWorkday(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday, Thursday)
}

func TestPermissionBits(t *testing.T) {
	var permission = 1
	t.Log(permission&Readable == Readable,
		permission&Writable == Writable,
		permission&Executable == Executable)

}
