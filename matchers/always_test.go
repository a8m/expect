package matchers_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/a8m/expect"
	"github.com/a8m/expect/matchers"
)

func TestFailsWhenMatcherFails(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	matcher := newMockMatcher()
	m := matchers.Always(matcher)
	matcher.MatchOutput.Ret0 <- fmt.Errorf("some-error")

	callCount := make(chan bool, 200)
	f := func() int {
		callCount <- true
		return 99
	}

	err := m.Match(f)
	expect(err).Not.To.Be.Nil()
	expect(callCount).To.Have.Len(1)

}

func TestPolls10TimesForSuccess(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	matcher := newMockMatcher()
	m := matchers.Always(matcher)

	callCount := make(chan bool, 200)
	f := func() int {
		callCount <- true
		matcher.MatchOutput.Ret0 <- nil
		return 101
	}
	m.Match(f)

	expect(callCount).To.Have.Len(10)
}

func TestPollsEach10msForSuccess(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	matcher := newMockMatcher()
	m := matchers.Always(matcher)

	var ts []int64
	f := func() int {
		ts = append(ts, time.Now().UnixNano())
		matcher.MatchOutput.Ret0 <- nil
		return 101
	}
	m.Match(f)

	for i := 0; i < len(ts)-1; i++ {
		expect(ts[i+1]-ts[i]).To.Be.Within(
			float64(10*time.Millisecond),
			float64(30*time.Millisecond),
		)
	}
}
