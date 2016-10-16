// +build go1.7

package expect_test

import (
	"testing"

	"github.com/a8m/expect"
)

func TestRun(t *testing.T) {
	mockT := newMockT()
	mockT.RunOutput.Succeeded <- true

	called := false
	ran := expect.Run(mockT, "Foo", func(expect expect.Expecter) {
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
	ran = expect.Run(mockT, "Bar", func(expect expect.Expecter) {
	})
	if ran {
		t.Error("expect.Run returned true when t.Run returned false")
	}
}
