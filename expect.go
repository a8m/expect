package expect

import (
	"testing"
)

type Expect struct {
	To  *to
	Not *not
}

type to struct {
	Be *Be
}

type not struct {
	To *to
}

func New(t *testing.T) func(v interface{}) *Expect {
	return func(v interface{}) *Expect {
		return &Expect{
			To:  &to{&Be{t, v, true}},
			Not: &not{&to{&Be{t, v, false}}},
		}
	}
}
