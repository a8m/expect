package expect

import (
	. "fmt"
	"reflect"
)

type Have struct {
	Else   *Else
	And    *Have
	t      T
	actual interface{}
	assert bool
}

func newHave(t T, e *Else, actual interface{}, assert bool) *Have {
	have := &Have{
		Else:   e,
		t:      t,
		actual: actual,
		assert: assert,
	}
	have.And = have
	return have
}

// Assert value to have length of the the given number
func (h *Have) Len(i int) *Have {
	msg := h.msg(Sprintf("length of %v", i))
	if l, ok := length(h.actual); ok {
		if l == i != h.assert {
			h.fail(2, msg)
		}
	} else {
		h.t.Fatal(invMsg("Array, Slice, Map or String"))
	}
	return h
}

// Assert value to have capacity of the given number
func (h *Have) Cap(i int) *Have {
	msg := h.msg(Sprint("capacity of %v", i))
	switch reflect.TypeOf(h.actual).Kind() {
	case reflect.Array, reflect.Slice, reflect.Chan:
		if reflect.ValueOf(h.actual).Cap() == i != h.assert {
			h.fail(2, msg)
		}
	default:
		h.t.Fatal(invMsg("Array, Slice or Chan"))
	}
	return h
}

// Assert `key` exists on the given Map, and has optional value.
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
				h.fail(2, msg)
			}
		} else {
			h.fail(2, msg)
		}
	default:
		h.t.Fatal(invMsg("Map"))
	}
	return h
}

// Assert `keys` exists on the given Map
func (h *Have) Keys(args ...interface{}) *Have {
	msg := h.msg(Sprintf("keys: %v", args))
	switch reflect.TypeOf(h.actual).Kind() {
	case reflect.Map:
		v := reflect.ValueOf(h.actual)
		for _, k := range args {
			vk := v.MapIndex(reflect.ValueOf(k))
			if vk.IsValid() != h.assert {
				h.fail(2, msg)
			}
		}
	default:
		h.t.Fatal(invMsg("Map"))
	}
	return h
}

// Assert `field` exist on the given Struct, and has optional value.
func (h *Have) Field(s string, args ...interface{}) *Have {
	// Test also value
	testVal := len(args) > 0
	msg := Sprintf("field: %v", s)
	if testVal {
		msg += Sprintf(" with value: %v", args[0])
	}
	msg = h.msg(msg)
	switch reflect.TypeOf(h.actual).Kind() {
	case reflect.Struct:
		v := reflect.ValueOf(h.actual)
		f := v.FieldByName(s)
		if (testVal && f.IsValid()) || f.IsValid() == h.assert {
			// Compare value
			if testVal && reflect.DeepEqual(f.Interface(), args[0]) != h.assert {
				h.fail(2, msg)
			}
		} else {
			h.fail(2, msg)
		}
	default:
		h.t.Fatal(invMsg("Struct"))
	}
	return h
}

// Assert `fields` exists on the given Struct
func (h *Have) Fields(args ...string) *Have {
	msg := h.msg(Sprintf("fields: %v", args))
	switch reflect.TypeOf(h.actual).Kind() {
	case reflect.Struct:
		v := reflect.ValueOf(h.actual)
		for _, f := range args {
			if v.FieldByName(f).IsValid() != h.assert {
				h.fail(2, msg)
			}
		}
	default:
		h.t.Fatal(invMsg("Struct"))
	}
	return h
}

// Assert `method` exist on the given struct/ptr
func (h *Have) Method(m string) *Have {
	msg := h.msg(Sprintf("method: %v", m))
	switch reflect.TypeOf(h.actual).Kind() {
	case reflect.Struct, reflect.Ptr:
		v := reflect.ValueOf(h.actual)
		if v.MethodByName(m).IsValid() != h.assert {
			h.fail(2, msg)
		}
	default:
		h.t.Fatal(invMsg("Struct or Ptr"))
	}
	return h
}

func (h *Have) fail(callers int, msg string) {
	h.Else.failed = true
	fail(h.t, callers, msg)
}

func (h *Have) msg(s string) string {
	return errMsg("to have")(h.actual, s, h.assert)
}
