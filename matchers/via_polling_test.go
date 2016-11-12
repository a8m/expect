package matchers_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/a8m/expect"
	"github.com/a8m/expect/matchers"
)

func TestViaPollingFailsPolls100Times(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	matcher := newMockMatcher()
	m := matchers.ViaPolling(matcher)

	callCount := make(chan bool, 200)
	f := func() int {
		callCount <- true
		matcher.MatchOutput.Ret0 <- fmt.Errorf("still wrong")
		return 99
	}
	m.Match(f)

	expect(callCount).To.Have.Len(100)
}

func TestViaPollingPollsEvery10ms(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	matcher := newMockMatcher()
	m := matchers.ViaPolling(matcher)

	var ts []int64
	f := func() int {
		ts = append(ts, time.Now().UnixNano())
		matcher.MatchOutput.Ret0 <- fmt.Errorf("still wrong")
		return 99
	}
	m.Match(f)

	for i := 0; i < len(ts)-1; i++ {
		expect(ts[i+1]-ts[i]).To.Be.Within(
			float64(10*time.Millisecond),
			float64(30*time.Millisecond),
		)
	}
}

func TestViaPollingStopsAfterSuccess(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	matcher := newMockMatcher()
	m := matchers.ViaPolling(matcher)

	callCount := make(chan bool, 200)
	f := func() int {
		callCount <- true
		matcher.MatchOutput.Ret0 <- nil
		return 101
	}
	m.Match(f)

	expect(callCount).To.Have.Len(1)
}

func TestViaPollingUsesGivenProperties(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	matcher := newMockMatcher()
	m := matchers.ViaPollingMatcher{
		Matcher:  matcher,
		Duration: time.Millisecond,
		Interval: 100 * time.Microsecond,
	}

	callCount := make(chan bool, 200)
	f := func() int {
		callCount <- true
		matcher.MatchOutput.Ret0 <- fmt.Errorf("still wrong")
		return 99
	}
	m.Match(f)

	expect(callCount).To.Have.Len(10)
}

func TestViaPollingPassesAlongChan(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	matcher := newMockMatcher()
	m := matchers.ViaPolling(matcher)
	matcher.MatchOutput.Ret0 <- nil

	c := make(chan int)
	m.Match(c)

	expect(matcher.MatchInput.Actual).To.Have.Len(1)
	expect(<-matcher.MatchInput.Actual).To.Equal(c)
}

func TestViaPollingFailsForNonChanOrFunc(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	matcher := newMockMatcher()
	m := matchers.ViaPolling(matcher)

	err := m.Match(101)
	expect(err).Not.To.Be.Nil()
}

func TestViaPollingFailsForFuncWithArgs(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	matcher := newMockMatcher()
	m := matchers.ViaPolling(matcher)

	err := m.Match(func(int) int { return 101 })
	expect(err).Not.To.Be.Nil()
}

func TestViaPollingFailsForFuncWithWrongReturns(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	matcher := newMockMatcher()
	m := matchers.ViaPolling(matcher)

	err := m.Match(func() (int, int) { return 101, 103 })
	expect(err).Not.To.Be.Nil()
}

func TestViaPollingFailsForSendOnlyChan(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	matcher := newMockMatcher()
	m := matchers.ViaPolling(matcher)

	err := m.Match(make(chan<- int))
	expect(err).Not.To.Be.Nil()
}

func TestViaPollingUsesChildMatcherErr(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	matcher := newMockMatcher()
	m := matchers.ViaPolling(matcher)

	f := func() int {
		matcher.MatchOutput.Ret0 <- fmt.Errorf("some-message")
		return 99
	}
	err := m.Match(f)
	expect(err).To.Equal(fmt.Errorf("some-message"))
}
