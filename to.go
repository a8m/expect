package expect

import (
	. "fmt"
	"reflect"
	"regexp"
	"strings"
)

type To struct {
	Be     *Be
	Have   *Have
	And    *To
	t      T
	actual interface{}
	assert bool
}

func NewTo(t T, actual interface{}, assert bool) *To {
	to := &To{
		t:      t,
		actual: actual,
		assert: assert,
	}
	to.Be = NewBe(t, actual, assert)
	to.Have = NewHave(t, actual, assert)
	to.And = to
	return to
}

// Assert that a string starts with `s`
func (t *To) StartWith(s string) *To {
	msg := t.msg(Sprintf("start with %v", s))
	if strings.HasPrefix(t.Str(), s) != t.assert {
		fail(t.t, 2, msg)
	}
	return t
}

// Assert that a string ends with `s`
func (t *To) EndWith(s string) *To {
	msg := t.msg(Sprintf("end with %v", s))
	if strings.HasSuffix(t.Str(), s) != t.assert {
		fail(t.t, 2, msg)
	}
	return t
}

// Assert that a string conatins `s`
func (t *To) Contains(s string) *To {
	msg := t.msg(Sprintf("contains %v", s))
	if strings.Contains(t.Str(), s) != t.assert {
		fail(t.t, 2, msg)
	}
	return t
}

// Assert whether a textual regular expression matches a string
func (t *To) Match(s string) *To {
	msg := t.msg(Sprintf("matches %v", s))
	matched, err := regexp.MatchString(s, t.Str())
	if err != nil {
		t.t.Fatal(err)
	}
	if matched != t.assert {
		fail(t.t, 2, msg)
	}
	return t
}

// Assert two values are equals(deeply)
func (t *To) Equal(exp interface{}) *To {
	msg := t.msg(Sprintf("equal to %v", exp))
	if reflect.DeepEqual(t.actual, exp) != t.assert {
		fail(t.t, 2, msg)
	}
	return t
}

// Assert func to panic
func (t *To) Panic(args ...interface{}) *To {
	testMsg := len(args) > 0
	switch t.actual.(type) {
	case func():
		fn := reflect.ValueOf(t.actual)
		if p, m := ifPanic(fn); p != t.assert || testMsg && args[0] == m != t.assert {
			if testMsg {
				m = args[0]
			}
			fail(t.t, 2, t.msg(Sprintf("panic: %v", m)))
		}
	default:
		t.t.Fatal(invMsg("func"))

	}
	return t
}

func ifPanic(f reflect.Value) (isPnc bool, msg interface{}) {
	func() {
		defer func() {
			if msg = recover(); msg != nil {
				isPnc = true
			}
		}()
		f.Call([]reflect.Value{})
	}()
	return
}

func (t *To) Str() (s string) {
	if s, ok := t.actual.(string); ok {
		return s
	}
	t.t.Fatal(invMsg("string"))
	return
}

func (t *To) msg(s string) string {
	return errMsg("to")(t.actual, s, t.assert)
}
