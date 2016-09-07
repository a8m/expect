package expect

type T interface {
	Errorf(format string, args ...interface{})
	Fatal(...interface{})
	FailNow()
}

type Expect struct {
	To  *To
	Not *Not
}

type Not struct {
	To *To
}

// Return new expect function with `To, To.Be, To.Have` assertions
func New(t T) func(v interface{}) *Expect {
	return func(v interface{}) *Expect {
		return &Expect{
			To:  newTo(t, v, true),
			Not: &Not{To: newTo(t, v, false)},
		}
	}
}
