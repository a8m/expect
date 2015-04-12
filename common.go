package expect

import (
	. "fmt"
	"reflect"
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

func invMsg(v string) (s string) {
	s = Sprintf("Invalid argument - expecting %v value.", v)
	return
}

func length(v interface{}) (int, bool) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice, reflect.Array, reflect.Map, reflect.String:
		return reflect.ValueOf(v).Len(), true
	default:
		return 0, false
	}
}
