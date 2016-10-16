// +build go1.7

package expect_test

import (
	"testing"

	"github.com/a8m/expect"
)

func TestRun(t *testing.T) {
	t.Run("Function", RunnerTests(expect.Run))
	e := expect.New(t)
	t.Run("Method", RunnerTests(e.Run))
}

func RunnerTests(run func(expect.T, string, func(*testing.T, expect.Expecter)) bool) func(t *testing.T) {
	return func(t *testing.T) {
		mockT := newMockT()
		mockT.RunOutput.Succeeded <- true

		called := false
		ran := run(mockT, "Foo", func(*testing.T, expect.Expecter) {
			called = true
		})
		if !ran {
			t.Error("expect.Run returned false when t.Run returned true")
		}

		select {
		case name := <-mockT.RunInput.Name:
			if name != "Foo" {
				t.Errorf(`%#v (actual) != "Foo" (expected)`, name)
			}
			test := <-mockT.RunInput.Test
			test(t)
			if !called {
				t.Errorf(`Calling test %#v did not call passed in expectation func`, test)
			}
		default:
			t.Error("expect.Run never called t.Run")
		}

		mockT.RunOutput.Succeeded <- false
		ran = run(mockT, "Bar", func(*testing.T, expect.Expecter) {
		})
		if ran {
			t.Error("expect.Run returned true when t.Run returned false")
		}
	}
}
