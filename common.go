package expect

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func errMsg(t string) func(act, exp interface{}, assert bool) string {
	return func(act, exp interface{}, assert bool) (s string) {
		not := "not "
		if assert {
			not = ""
		}
		s = fmt.Sprintf("Expect %v %v%v %v", act, not, t, exp)
		return
	}
}

func invMsg(v string) (s string) {
	s = fmt.Sprintf("Invalid argument - expecting %v value.", v)
	return
}

func length(v interface{}) (int, bool) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice, reflect.Array, reflect.Map, reflect.String, reflect.Chan:
		return reflect.ValueOf(v).Len(), true
	default:
		return 0, false
	}
}

// fail creates a message including the stack trace returned from
// runtime.Caller
func fail(t *testing.T, callers int, msg string) {
	stack := stackPrefix(fullStack(callers + 1))
	// It makes sense to have the most recent call closest to the
	// message.
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}
	prefix := strings.Join(stack, "\n\r\t")
	t.Errorf("\r\t%s\n%s", prefix, msg)
}

// fullStack gets the full call stack, skipping the passed in number
// of callers.
func fullStack(skip int) []uintptr {
	var stack []uintptr
	next := make([]uintptr, 10)
	for {
		added := runtime.Callers(skip, next)
		stack = append(stack, next[:added]...)
		if added < len(next) {
			break
		}
		skip += added
	}
	return stack
}
