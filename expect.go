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
		// Have, Not.Have
		var have, nhave Have
		have = Have{t, &have, v, true}
		nhave = Have{t, &nhave, v, false}
		// To, Not.To
		var to, nto To
		to = To{t, &be, &have, &to, v, true}
		nto = To{t, &nbe, &nhave, &nto, v, false}
		return &Expect{
			To:  &to,
			Not: &not{&nto},
		}
	}
}
