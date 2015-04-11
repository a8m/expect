package expect

import (
	"testing"
)

type Expect struct {
	To  *To
	Not *not
}

type not struct {
	To *To
}

func New(t *testing.T) func(v interface{}) *Expect {
	return func(v interface{}) *Expect {
		return &Expect{
			To:  &To{t, &Be{t, v, true}, v, true},
			Not: &not{&To{t, &Be{t, v, false}, v, false}},
		}
	}
}
