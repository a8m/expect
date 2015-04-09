package should

import (
	"fmt"
	"testing"
)

func New(t *testing.T) func(v interface{}) {
	return func(v interface{}) {
		fmt.Println(v, t)
		t.Error("fail")
	}
}
