package expect

import "testing"

// TODO(Ariel): Create mock that implement TB interface
// and stub `Error` and `Fatal`

func TestAbove(t *testing.T) {
	expect := New(t)
	expect(10).To.Be.Above(0)
	expect(10).Not.To.Be.Above(20)
}

func TestBelow(t *testing.T) {
	expect := New(t)
	expect(10).To.Be.Below(20)
	expect(10).Not.To.Be.Below(0)
}

func TestBeChaining(t *testing.T) {
	expect := New(t)
	expect(10).To.Be.Above(0).And.Below(20)
	expect(10).Not.To.Be.Above(20).And.Below(0)
}
