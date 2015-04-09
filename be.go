package should

import (
	"fmt"
	"testing"
)

type Be struct {
	*testing.T
	Actual interface{}
}

func (b *Be) Above(e interface{}) {
	if i, ok := e.(int); ok {

	}
	b.Error("foo")
}
