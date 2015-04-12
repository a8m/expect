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

func TestEmpty(t *testing.T) {
	expect := New(t)
	expect("").To.Be.Empty()
	expect([]int{1, 2, 3}).Not.To.Be.Empty()
	expect(make(map[string]int)).To.Be.Empty()
	expect([2]int{}).Not.To.Be.Empty()
	expect([]byte{}).To.Be.Empty()
}

func TestOk(t *testing.T) {
	expect := New(t)
	expect("").Not.To.Be.Ok()
	expect(false).Not.To.Be.Ok()
	expect(0).Not.To.Be.Ok()
	expect(nil).Not.To.Be.Ok()

	expect("foo").To.Be.Ok()
	expect(1).To.Be.Ok()
	expect(true).To.Be.Ok()
	expect(struct{}{}).To.Be.Ok()
	expect([]int{}).To.Be.Ok()
}

func TestString(t *testing.T) {
	expect := New(t)
	expect("").To.Be.String()
	expect(1).Not.To.Be.String()
}

func TestBeChaining(t *testing.T) {
	expect := New(t)
	expect(10).To.Be.Above(0).And.Below(20)
	expect(10).Not.To.Be.Above(20).And.Below(0)
}
