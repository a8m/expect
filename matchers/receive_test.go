package matchers_test

import (
	"fmt"
	"testing"

	"github.com/a8m/expect"
	"github.com/a8m/expect/matchers"
)

func TestReceiveSucceedsForABufferedChannel(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	c := make(chan bool, 1)
	m := matchers.Receive()

	err := m.Match(c)
	expect(err).Not.To.Be.Nil()

	c <- true
	err = m.Match(c)
	expect(err).To.Be.Nil()
}

func TestReceiveToSetsArg(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	c := make(chan int, 1)

	var value int
	m := matchers.ReceiveTo(&value)

	err := m.Match(c)
	expect(err).Not.To.Be.Nil()
	expect(value).To.Equal(0)

	c <- 17

	err = m.Match(c)
	expect(err).To.Be.Nil()
	expect(value).To.Equal(17)
}

func TestReceiveFailsNotReadableChan(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	c := make(chan int, 10)
	m := matchers.Receive()

	err := m.Match(101)
	expect(err).Not.To.Be.Nil()

	err = m.Match(chan<- int(c))
	expect(err).Not.To.Be.Nil()
}

func TestReceiveToFailsInvalidType(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	c := make(chan int, 10)
	c <- 17
	c <- 17

	var value float64
	m := matchers.ReceiveTo(&value)

	err := m.Match(c)
	expect(err).Not.To.Be.Nil()

	var nonPointer int
	m = matchers.ReceiveTo(nonPointer)

	err = m.Match(c)
	fmt.Println(err)
	expect(err).Not.To.Be.Nil()
}

func TestReceiveFailsForClosedChannel(t *testing.T) {
	t.Parallel()
	expect := expect.New(t)
	c := make(chan bool, 1)
	m := matchers.Receive()
	c <- true

	err := m.Match(c)
	expect(err).To.Be.Nil()

	close(c)
	err = m.Match(c)
	expect(err).Not.To.Be.Nil()
}
