package expect

import (
	. "fmt"
	"reflect"
)

type Be struct {
	Else   *Else
	And    *Be
	t      T
	actual interface{}
	assert bool
}

func NewBe(t T, actual interface{}, assert bool) *Be {
	be := &Be{
		Else:   NewElse(t),
		t:      t,
		actual: actual,
		assert: assert,
	}
	be.And = be
	return be
}

// Assert numeric value above the given value (> n)
func (b *Be) Above(e float64) *Be {
	msg := b.msg(Sprintf("above %v", e))
	if b.Num() > e != b.assert {
		b.fail(2, msg)
	}
	return b
}

// Assert numeric value below the given value (< n)
func (b *Be) Below(e float64) *Be {
	msg := b.msg(Sprintf("below %v", e))
	if b.Num() < e != b.assert {
		b.fail(2, msg)
	}
	return b
}

// Assert inclusive numeric range (<= to and >= from)
func (b *Be) Within(from, to float64) *Be {
	msg := b.msg(Sprintf("between range %v <= x <= %v", from, to))
	x := b.Num()
	if x <= to && x >= from != b.assert {
		b.fail(2, msg)
	}
	return b
}

// Assert given value is empty, Array, Slice, Map or String
func (b *Be) Empty() *Be {
	msg := b.msg("empty")
	if i, ok := length(b.actual); ok {
		if i == 0 != b.assert {
			b.fail(2, msg)
		}
	} else {
		b.t.Fatal(invMsg("Array, Slice, Map or String"))
	}
	return b
}

// Assert if the given value is truthy(i.e: not "", nil, 0, false)
func (b *Be) Ok() *Be {
	msg := b.msg("ok")
	var exp bool
	switch b.actual.(type) {
	case int, int8, int32, int64, uint, uint8, uint32, uint64, float32, float64:
		exp = b.actual != 0
	case string:
		exp = b.actual != ""
	case bool:
		exp = b.actual != false // TODO(Ariel): without the `!= false`, it's ask for type assertion
	default:
		exp = b.actual != nil
	}
	if exp != b.assert {
		b.fail(2, msg)
	}
	return b
}

// Assert given value is type of string
func (b *Be) String() *Be {
	msg := b.msg("string")
	if _, ok := b.actual.(string); ok != b.assert {
		b.fail(2, msg)
	}
	return b
}

// Assert given value is type of int
func (b *Be) Int() *Be {
	msg := b.msg("int")
	if _, ok := b.actual.(int); ok != b.assert {
		b.fail(2, msg)
	}
	return b
}

// Assert given value is type of float32/64
func (b *Be) Float() *Be {
	msg := b.msg("float")
	exp := false
	switch b.actual.(type) {
	case float32, float64:
		exp = true
	}
	if exp != b.assert {
		b.fail(2, msg)
	}
	return b
}

// Assert given value is type of boolean
func (b *Be) Bool() *Be {
	msg := b.msg("boolean")
	if _, ok := b.actual.(bool); ok != b.assert {
		b.fail(2, msg)
	}
	return b
}

// Assert given value is type of map
func (b *Be) Map() *Be {
	msg := b.msg("map")
	if reflect.TypeOf(b.actual).Kind() == reflect.Map != b.assert {
		b.fail(2, msg)
	}
	return b
}

// Assert given value is type of array
func (b *Be) Array() *Be {
	msg := b.msg("array")
	if reflect.TypeOf(b.actual).Kind() == reflect.Array != b.assert {
		b.fail(2, msg)
	}
	return b
}

// Assert given value is type of slice
func (b *Be) Slice() *Be {
	msg := b.msg("slice")
	if reflect.TypeOf(b.actual).Kind() == reflect.Slice != b.assert {
		b.fail(2, msg)
	}
	return b
}

// Assert given value is type of channel
func (b *Be) Chan() *Be {
	msg := b.msg("channel")
	if reflect.TypeOf(b.actual).Kind() == reflect.Chan != b.assert {
		b.fail(2, msg)
	}
	return b
}

// Assert given value is type of struct
func (b *Be) Struct() *Be {
	msg := b.msg("struct")
	if reflect.TypeOf(b.actual).Kind() == reflect.Struct != b.assert {
		b.fail(2, msg)
	}
	return b
}

// Assert given value is type of pointer
func (b *Be) Ptr() *Be {
	msg := b.msg("pointer")
	if reflect.TypeOf(b.actual).Kind() == reflect.Ptr != b.assert {
		b.fail(2, msg)
	}
	return b
}

// Assert given value is nil
func (b *Be) Nil() *Be {
	msg := b.msg("nil")
	if b.actual == nil != b.assert {
		b.fail(2, msg)
	}
	return b
}

// Assert given value is type of the given string
func (b *Be) Type(s string) *Be {
	msg := b.msg(Sprintf("type %v", s))
	if reflect.TypeOf(b.actual).Name() == s != b.assert {
		b.fail(2, msg)
	}
	return b
}

func (b *Be) fail(callers int, msg string) {
	b.Else.failed = true
	fail(b.t, callers, msg)
}

func (b *Be) msg(s string) string {
	return errMsg("to be")(b.actual, s, b.assert)
}

func (b *Be) Num() float64 {
	rv := reflect.ValueOf(b.actual)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(rv.Int())
	case reflect.Uint, reflect.Uintptr, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(rv.Uint())
	case reflect.Float32, reflect.Float64:
		return float64(rv.Float())
	default:
		b.t.Fatal(invMsg("numeric"))
		return 0
	}
}
