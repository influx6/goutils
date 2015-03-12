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

func TestString(t *testing.T) {
	var item interface{} = 1000
	morph := NewStringMorpher()
	val := morph.Morph(item)

	if val != "1000" {
		t.Fatalf("expecting `1000` but got incorrect value", val, morph, item)
	}
}

func TestMap(t *testing.T) {
	store := NewMap()

	store.Set("day", "build")

	if !store.Has("day") {
		t.Fatalf("expecting `true` but got false", store, "day")
	}

	if !store.HasMatch("day", "build") {
		t.Fatalf("expecting `true` but got false", store, "day")
	}

	store.Remove("day")

	if store.Has("day") {
		t.Fatalf("expecting `false` but got true", store, "day")
	}
}
