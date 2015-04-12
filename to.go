package expect

import (
	. "fmt"
	"regexp"
	"strings"
	"testing"
)

type To struct {
	*testing.T
	Be     *Be
	Have   *Have
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

// Assert  whether a textual regular expression matches a string
func (t *To) Match(s string) *To {
	msg := t.msg(Sprintf("matches %v", s))
	matched, err := regexp.MatchString(s, t.Str())
	if err != nil {
		t.Fatal(err)
	}
	if matched != t.assert {
		t.Error(msg)
	}
	return t
}

func (t *To) Str() (s string) {
	if s, ok := t.actual.(string); ok {
		return s
	}
	t.Fatal(invMsg("string"))
	return
}

func (t *To) msg(s string) string {
	return errMsg("to")(t.actual, s, t.assert)
}
