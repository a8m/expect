package expect

import (
	. "fmt"
	"reflect"
	"testing"
)

type Be struct {
	*testing.T
	And    *Be
	actual interface{}
	assert bool
}

// Assert numeric value above the given value (> n)
func (b *Be) Above(e float64) *Be {
	msg := b.msg(Sprintf("above %v", e))
	if b.Num() > e != b.assert {
		b.Error(msg)
	}
	return b
}

// Assert numeric value below the given value (< n)
func (b *Be) Below(e float64) *Be {
	msg := b.msg(Sprintf("below %v", e))
	if b.Num() < e != b.assert {
		b.Error(msg)
	}
	return b
}

func (b *Be) msg(s string) string {
	not := "not "
	if b.assert {
		not = ""
	}
	return Sprintf("Expect %v %vto be %v", b.actual, not, s)
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
		b.Fatal("Invalid argument - expecting numeric value.")
		return 0
	}
}
