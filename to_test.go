package expect

import "testing"

// TODO(Ariel): Create mock that implement TB interface
// and stub `Error` and `Fatal`

func TestStartWith(t *testing.T) {
	expect := New(t)
	expect("foo").To.StartWith("f")
	expect("foo").Not.To.StartWith("bar")
}
