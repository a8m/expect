// +build go1.7

package expect

import "testing"

// Run acts like t.Run, but performs the `expect.New(t)` step for
// you, passing in the resulting Expecter.
func Run(t T, name string, expectation func(*testing.T, Expecter)) bool {
	return t.Run(name, func(t *testing.T) {
		expect := New(t)
		expectation(t, expect)
	})
}

// Run is a convenience method for use when the expect package is
// shadowed, e.g. `expect := expect.New(t)`.
func (Expecter) Run(t T, name string, expectation func(*testing.T, Expecter)) bool {
	return Run(t, name, expectation)
}
