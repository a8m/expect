package expect

type Else struct {
	t      T
	failed bool
}

func newElse(t T) *Else {
	return &Else{
		t: t,
	}
}

func (e *Else) FailNow() {
	if e.failed {
		e.t.FailNow()
	}
}
