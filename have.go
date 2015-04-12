package expect

import (
	. "fmt"
	"reflect"
	"testing"
)

type Have struct {
	*testing.T
	And    *Have
	actual interface{}
	assert bool
}

// Assert value to have length of the the given number
func (h *Have) Len(i int) *Have {
	msg := h.msg(Sprintf("length of %v", i))
	if l, ok := length(h.actual); ok {
		if l == i != h.assert {
			h.Error(msg)
		}
	} else {
		h.Fatal(invMsg("Array, Slice, Map or String"))
	}
	return h
}

func (h *Have) Key(args ...interface{}) *Have {
	// Test also value
	testVal := len(args) > 1
	msg := Sprintf("key: %v", args[0])
	if testVal {
		msg += Sprintf(" with value: %v", args[1])
	}
	msg = h.msg(msg)
	switch reflect.TypeOf(h.actual).Kind() {
	case reflect.Map:
		v := reflect.ValueOf(h.actual)
		k := v.MapIndex(reflect.ValueOf(args[0]))
		if (testVal && k.IsValid()) || k.IsValid() == h.assert {
			// Compare value
			if testVal && reflect.DeepEqual(k.Interface(), args[1]) != h.assert {
				h.Error(msg)
			}
		} else {
			h.Error(msg)
		}
	default:
		h.Fatal(invMsg("Map"))
	}
	return h
}

func (h *Have) msg(s string) string {
	return errMsg("to have")(h.actual, s, h.assert)
}
