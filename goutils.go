package goutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
)

var (
	MorphString = NewStringMorpher()
	MorphByte   = NewByteMorpher()
)

// Equal is a helper for comparing value equality, following these rules:
//  - Values with equivalent types are compared with reflect.DeepEqual
//  - int, uint, and float values are compared without regard to the type width.
//    for example, Equal(int32(5), int64(5)) == true
//  - strings and byte slices are converted to strings before comparison.
//  - else, return false.

func Equal(a, b interface{}) bool {
	if reflect.TypeOf(a) == reflect.TypeOf(b) {
		return reflect.DeepEqual(a, b)
	}
	switch a.(type) {
	case int, int8, int16, int32, int64:
		switch b.(type) {
		case int, int8, int16, int32, int64:
			return reflect.ValueOf(a).Int() == reflect.ValueOf(b).Int()
		}
	case uint, uint8, uint16, uint32, uint64:
		switch b.(type) {
		case uint, uint8, uint16, uint32, uint64:
			return reflect.ValueOf(a).Uint() == reflect.ValueOf(b).Uint()
		}
	case float32, float64:
		switch b.(type) {
		case float32, float64:
			return reflect.ValueOf(a).Float() == reflect.ValueOf(b).Float()
		}
	case string:
		switch b.(type) {
		case []byte:
			return a.(string) == string(b.([]byte))
		}
	case []byte:
		switch b.(type) {
		case string:
			return b.(string) == string(a.([]byte))
		}
	}
	return false
}

type Map struct {
	internal map[interface{}]interface{}
}

func (m *Map) HasMatch(key, value interface{}) bool {
	k, ok := m.internal[key]

	if ok {
		return k == value
	}

	return false
}

func (m *Map) Clone(src *Map) {
	for k, v := range src.Map() {
		m.Set(k, v)
	}
}

func (m *Map) Copy(src map[interface{}]interface{}) {
	for k, v := range src {
		m.Set(k, v)
	}
}

func (m *Map) Has(key interface{}) bool {
	_, ok := m.internal[key]
	return ok
}

func (m *Map) Get(key interface{}) interface{} {
	return m.internal[key]
}

func (m *Map) Set(key, value interface{}) {
	m.internal[key] = value
}

func (m *Map) Remove(key interface{}) {
	delete(m.internal, key)
}

func (m *Map) Map() map[interface{}]interface{} {
	return m.internal
}

func NewMap() *Map {
	return &Map{make(map[interface{}]interface{})}
}

type TypeCallers struct {
	Int     func(int)
	UInt    func(uint)
	Int8    func(int8)
	UInt8   func(uint8)
	Int16   func(int16)
	UInt16  func(uint16)
	Int32   func(int32)
	UInt32  func(uint32)
	Int64   func(int64)
	UInt64  func(uint64)
	String  func(string)
	Byte    func(byte)
	Bytes   func([]byte)
	Float64 func(float64)
	Float32 func(float32)
	Unknown func(interface{})
}

type ByteMorpher struct {
	stringmorph *StringMorpher
}

func (b *ByteMorpher) Morph(a interface{}) []byte {
	val := b.stringmorph.Morph(a)
	return []byte(val)
}

func NewByteMorpher() *ByteMorpher {
	return &ByteMorpher{NewStringMorpher()}
}

type String struct {
	Value string
}

func (s *String) String() string {
	return s.Value
}

func NewString(a string) *String {
	return &String{a}
}

type StringMorpher struct {
	buffer    *String
	converter *TypeCallers
}

func (s *StringMorpher) Morph(a interface{}) string {
	OnType(a, s.converter)
	return s.buffer.Value
}

func NewStringMorpher() *StringMorpher {
	val := &String{""}
	conv := NewStringConverter(val)
	return &StringMorpher{val, conv}
}

func NewStringConverter(val *String) *TypeCallers {
	return &TypeCallers{
		func(item int) {
			val.Value = fmt.Sprint(item)
		},
		func(item uint) {
			val.Value = fmt.Sprint(item)
		},
		func(item int8) {
			val.Value = fmt.Sprint(item)
		},
		func(item uint8) {
			val.Value = fmt.Sprint(item)
		},
		func(item int16) {
			val.Value = fmt.Sprint(item)
		},
		func(item uint16) {
			val.Value = fmt.Sprint(item)
		},
		func(item int32) {
			val.Value = fmt.Sprint(item)
		},
		func(item uint32) {
			val.Value = fmt.Sprint(item)
		},
		func(item int64) {
			val.Value = fmt.Sprint(item)
		},
		func(item uint64) {
			val.Value = fmt.Sprint(item)
		},
		func(item string) {
			val.Value = item
		},
		func(item byte) {
			val.Value = fmt.Sprint(item)
		},
		func(item []byte) {
			val.Value = string(item)
		},
		func(item float64) {
			val.Value = fmt.Sprint(item)
		},
		func(item float32) {
			val.Value = fmt.Sprint(item)
		},
		func(item interface{}) {
			conv, err := json.Marshal(item)

			if err != nil {
				log.Println(err)
				return
			}

			val.Value = string(conv)
		},
	}
}

func OnType(a interface{}, caller *TypeCallers) {
	if val, err := ByteListMorph(a); err == nil {
		if caller.Bytes != nil {
			caller.Bytes(val)
		}
		return
	}

	if val, err := ByteMorph(a); err == nil {
		if caller.Byte != nil {
			caller.Byte(val)
		}
		return
	}
	if val, err := StringMorph(a); err == nil {
		if caller.String != nil {
			caller.String(val)
		}
		return
	}
	if val, err := Float32Morph(a); err == nil {
		if caller.Float32 != nil {
			caller.Float32(val)
		}
		return
	}
	if val, err := Float64Morph(a); err == nil {
		if caller.Float64 != nil {
			caller.Float64(val)
		}
		return
	}
	if val, err := Int64Morph(a); err == nil {
		if caller.Int64 != nil {
			caller.Int64(val)
		}
		return
	}
	if val, err := UInt64Morph(a); err == nil {
		if caller.UInt64 != nil {
			caller.UInt64(val)
		}
		return
	}
	if val, err := Int32Morph(a); err == nil {
		if caller.Int32 != nil {
			caller.Int32(val)
		}
		return
	}
	if val, err := UInt32Morph(a); err == nil {
		if caller.UInt32 != nil {
			caller.UInt32(val)
		}
		return
	}
	if val, err := Int16Morph(a); err == nil {
		if caller.Int16 != nil {
			caller.Int16(val)
		}
		return
	}
	if val, err := UInt16Morph(a); err == nil {
		if caller.UInt16 != nil {
			caller.UInt16(val)
		}
		return
	}
	if val, err := Int8Morph(a); err == nil {
		if caller.Int8 != nil {
			caller.Int8(val)
		}
		return
	}
	if val, err := UInt8Morph(a); err == nil {
		if caller.UInt8 != nil {
			caller.UInt8(val)
		}
		return
	}
	if val, err := IntMorph(a); err == nil {
		if caller.Int != nil {
			caller.Int(val)
		}
		return
	}
	if val, err := UIntMorph(a); err == nil {
		if caller.UInt != nil {
			caller.UInt(val)
		}
		return
	}

	if caller.Unknown != nil {
		caller.Unknown(a)
	}
}

func IsBasicType(a interface{}) bool {
	if _, err := ByteListMorph(a); err != nil {
		return false
	}
	if _, err := ByteMorph(a); err != nil {
		return false
	}
	if _, err := StringMorph(a); err != nil {
		return false
	}
	if _, err := Float32Morph(a); err != nil {
		return false
	}
	if _, err := Float64Morph(a); err != nil {
		return false
	}
	if _, err := Int64Morph(a); err != nil {
		return false
	}
	if _, err := UInt64Morph(a); err != nil {
		return false
	}
	if _, err := Int32Morph(a); err != nil {
		return false
	}
	if _, err := UInt32Morph(a); err != nil {
		return false
	}
	if _, err := Int16Morph(a); err != nil {
		return false
	}
	if _, err := UInt16Morph(a); err != nil {
		return false
	}
	if _, err := Int8Morph(a); err != nil {
		return false
	}
	if _, err := UInt8Morph(a); err != nil {
		return false
	}
	if _, err := IntMorph(a); err != nil {
		return false
	}
	if _, err := UIntMorph(a); err != nil {
		return false
	}

	return true
}

func ByteListMorph(a interface{}) ([]byte, error) {
	m, ok := a.([]byte)

	if !ok {
		return nil, errors.New("Not a string")
	}

	return m, nil
}

func ByteMorph(a interface{}) (byte, error) {
	m, ok := a.(byte)

	if !ok {
		return *new(byte), errors.New("Not a string")
	}

	return m, nil
}

func StringMorph(a interface{}) (string, error) {
	m, ok := a.(string)

	if !ok {
		return *new(string), errors.New("Not a string")
	}

	return m, nil
}

func Float32Morph(a interface{}) (float32, error) {
	m, ok := a.(float32)

	if !ok {
		return *new(float32), errors.New("Not a float32")
	}

	return m, nil
}

func Float64Morph(a interface{}) (float64, error) {
	m, ok := a.(float64)

	if !ok {
		return *new(float64), errors.New("Not a float64")
	}

	return m, nil
}

func UInt16Morph(a interface{}) (uint16, error) {
	m, ok := a.(uint16)

	if !ok {
		return *new(uint16), errors.New("Not a int16")
	}

	return m, nil
}

func UInt32Morph(a interface{}) (uint32, error) {
	m, ok := a.(uint32)

	if !ok {
		return *new(uint32), errors.New("Not a uint32")
	}

	return m, nil
}

func UInt64Morph(a interface{}) (uint64, error) {
	m, ok := a.(uint64)

	if !ok {
		return *new(uint64), errors.New("Not a uint64")
	}

	return m, nil
}

func UIntMorph(a interface{}) (uint, error) {
	m, ok := a.(uint)

	if !ok {
		return *new(uint), errors.New("Not a uint")
	}

	return m, nil
}

func UInt8Morph(a interface{}) (uint8, error) {
	m, ok := a.(uint8)

	if !ok {
		return *new(uint8), errors.New("Not a uint8")
	}

	return m, nil
}

func Int16Morph(a interface{}) (int16, error) {
	m, ok := a.(int16)

	if !ok {
		return *new(int16), errors.New("Not a int16")
	}

	return m, nil
}

func Int32Morph(a interface{}) (int32, error) {
	m, ok := a.(int32)

	if !ok {
		return *new(int32), errors.New("Not a int32")
	}

	return m, nil
}

func Int64Morph(a interface{}) (int64, error) {
	m, ok := a.(int64)

	if !ok {
		return *new(int64), errors.New("Not a int64")
	}

	return m, nil
}

func IntMorph(a interface{}) (int, error) {
	m, ok := a.(int)

	if !ok {
		return *new(int), errors.New("Not a int")
	}

	return m, nil
}

func Int8Morph(a interface{}) (int8, error) {
	m, ok := a.(int8)

	if !ok {
		return *new(int8), errors.New("Not a int8")
	}

	return m, nil
}
