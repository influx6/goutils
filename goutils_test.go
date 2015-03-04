package goutils

import (
	"testing"
)

func TestInt(t *testing.T) {
	var item interface{}
	item = 70

	real, err := IntMorph(item)

	if err != nil {
		t.Fatal("error converting interface into int type", err, item, real)
	}

}

func TestBytes(t *testing.T) {
	var item interface{}
	item = []byte("70")

	real, err := ByteListMorph(item)

	if err != nil {
		t.Fatal("error converting interface into int []byte", err, item, real)
	}

}

func TestCallers(t *testing.T) {
	var item interface{}
	item = 70

	stack := new(TypeCallers)

	stack.Int = func(data int) {
		t.Logf("Passed for int call %d %d", data, item)
	}

	OnType(item, stack)
}
