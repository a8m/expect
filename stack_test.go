package expect_test

import (
	"strings"
	"testing"

	"github.com/a8m/expect"
)

func TestStackTrace(t *testing.T) {
	mockT := newMockT()
	expect := expect.New(mockT)
	expect("foo").To.Equal("bar")

	var args []interface{}
	select {
	case args = <-mockT.ErrorfInput.Args:
	default:
		t.Fatal("Errorf was never called for a failing expectation")
	}
	if len(args) != 2 {
		t.Fatalf("Wrong number of arguments (expected 2): %d", args)
	}
	stack, ok := args[0].(string)
	if !ok {
		t.Fatalf("Expected %#v to be a string")
	}
	if !strings.Contains(stack, "stack_test.go") {
		t.Errorf(`Expected "%s" to contain "stack_test.go"`, stack)
	}
}
