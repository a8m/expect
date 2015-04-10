package expect

import (
	. "fmt"
	"testing"
)

type Be struct {
	*testing.T
	actual interface{}
	assert bool
}

func (b *Be) Above(e int) {
	msg := b.msg(Sprintf("above %v", e))
	if b.Int() > e != b.assert {
		b.Error(msg)
	}
}

func (b *Be) Below(e int) {
	msg := b.msg(Sprintf("below %v", e))
	if b.Int() < e != b.assert {
		b.Error(msg)
	}
}

func (b *Be) msg(s string) string {
	not := "not "
	if b.assert {
		not = ""
	}
	return Sprintf("Expect %v %vto be %v", b.actual, not, s)
}

func (b *Be) Int() int {
	if i, ok := b.actual.(int); ok {
		return i
	}
	b.Fatal("Invalid argument, expect to int")
	return 0
}
