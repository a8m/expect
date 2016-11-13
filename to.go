package expect

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/a8m/expect/matchers"
)

type To struct {
	Be     *Be
	Have   *Have
	Else   *Else
	And    *To
	t      T
	actual interface{}
	assert bool
}

func newTo(t T, actual interface{}, assert bool) *To {
	to := &To{
		t:      t,
		actual: actual,
		assert: assert,
	}
	to.Else = newElse(t)
	to.Be = newBe(t, to.Else, actual, assert)
	to.Have = newHave(t, to.Else, actual, assert)
	to.And = to
	return to
}

// Assert that a string starts with `s`
func (t *To) StartWith(s string) *To {
	msg := t.msg(fmt.Sprintf("start with %v", s))
	if strings.HasPrefix(t.Str(), s) != t.assert {
		t.fail(2, msg)
	}
	return t
}

// Assert that a string ends with `s`
func (t *To) EndWith(s string) *To {
	msg := t.msg(fmt.Sprintf("end with %v", s))
	if strings.HasSuffix(t.Str(), s) != t.assert {
		t.fail(2, msg)
	}
	return t
}

// Assert that a string conatins `s`
func (t *To) Contains(s string) *To {
	msg := t.msg(fmt.Sprintf("contains %v", s))
	if strings.Contains(t.Str(), s) != t.assert {
		t.fail(2, msg)
	}
	return t
}

// Assert whether a textual regular expression matches a string
func (t *To) Match(s string) *To {
	msg := t.msg(fmt.Sprintf("matches %v", s))
	matched, err := regexp.MatchString(s, t.Str())
	if err != nil {
		t.t.Fatal(err)
	}
	if matched != t.assert {
		t.fail(2, msg)
	}
	return t
}

// Assert two values are equals(deeply)
func (t *To) Equal(exp interface{}) *To {
	msg := t.msg(fmt.Sprintf("equal to %v", exp))
	if reflect.DeepEqual(t.actual, exp) != t.assert {
		t.fail(2, msg)
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
			t.fail(2, t.msg(fmt.Sprintf("panic: %v", m)))
		}
	default:
		t.t.Fatal(invMsg("func"))

	}
	return t
}

func (t *To) Pass(matcher Matcher) *To {
	err := matcher.Match(t.actual)
	switch t.assert {
	case true:
		if err != nil {
			t.fail(2, t.msg(err.Error()))
		}
	case false:
		if err == nil {
			t.fail(2, t.msg(fmt.Sprintf("match %#v", matcher)))
		}
	}
	return t
}

// Assert that a value can be read from a channel
func (t *To) Receive() *To {
	var value interface{}
	err := matchers.ReceiveTo(&value).Match(t.actual)
	if t.assert && err != nil {
		t.fail(2, t.msg(err.Error()))
	} else if !t.assert && err == nil {
		t.fail(2, t.msg("not to receive"))
	}

	return newTo(t.t, value, t.assert)
}

func (t *To) fail(callers int, msg string) {
	fail(t.t, callers+1, msg)
	t.Else.failed = true
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
