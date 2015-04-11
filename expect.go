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
		// Be, Not.Be
		var be, nbe Be
		be = Be{t, &be, v, true}
		nbe = Be{t, &nbe, v, false}
		// To, Not.To
		var to, nto To
		to = To{t, &be, &to, v, true}
		nto = To{t, &nbe, &nto, v, false}
		return &Expect{
			To:  &to,
			Not: &not{&nto},
		}
	}
}
