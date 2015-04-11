package expect

import (
	. "fmt"
	"strings"
	"testing"
)

type To struct {
	*testing.T
	Be     *Be
	And    *To
	actual interface{}
	assert bool
}

// Assert that a string starts with `s`
func (t *To) StartWith(s string) *To {
	msg := t.msg(Sprintf("start with %v", s))
	if strings.HasPrefix(t.Str(), s) != t.assert {
		t.Error(msg)
	}
	return t
}

// Assert that a string ends with `s`
func (t *To) EndWith(s string) *To {
	msg := t.msg(Sprintf("end with %v", s))
	if strings.HasSuffix(t.Str(), s) != t.assert {
		t.Error(msg)
	}
	return t
}

// Assert that a string conatins `s`
func (t *To) Contains(s string) *To {
	msg := t.msg(Sprintf("contains %v", s))
	if strings.Contains(t.Str(), s) != t.assert {
		t.Error(msg)
	}
	return t
}

func (t *To) Str() (s string) {
	if s, ok := t.actual.(string); ok {
		return s
	}
	t.Error("Ivalid argument - expecting string value")
	return
}

func (t *To) msg(s string) string {
	not := "not "
	if t.assert {
		not = ""
	}
	return Sprintf("Expect %v %vto %v", t.actual, not, s)
}
