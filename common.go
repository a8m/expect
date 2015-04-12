package expect

import (
	. "fmt"
)

func errMsg(t string) func(act, exp interface{}, assert bool) string {
	return func(act, exp interface{}, assert bool) (s string) {
		not := "not "
		if assert {
			not = ""
		}
		s = Sprintf("Expect %v %v%v %v", act, not, t, exp)
		return
	}
}
