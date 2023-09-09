package types

import (
	"math"
	"testing"
)

func TestGoTypes(t *testing.T) {

}

type Myint int32

func TestImplicitTypeConversion(t *testing.T) {
	var a int = 1
	var b int32
	b = int32(a)
	t.Log(a, b)
	var c Myint
	c = Myint(b)
	t.Log(b, c)
	t.Log(math.MaxInt64, math.MaxFloat64, math.MaxUint32)
}

func TestPointer(t *testing.T) {
	a := 1
	var aPtr = &a
	//aPtr = aPtr + 1
	t.Log(a, aPtr, &aPtr)
	t.Logf("%T %T %T", a, aPtr, &aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
