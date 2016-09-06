package expect_test

import (
	"testing"

	"github.com/a8m/expect"
)

// TODO(Ariel): Create mock that implement TB interface
// and stub `Error` and `Fatal`

func TestLen(t *testing.T) {
	expect := expect.New(t)
	expect("foo").To.Have.Len(3)
	m := map[string]int{}
	expect(m).To.Have.Len(0)
	expect(m).Not.To.Have.Len(1)
	s := []string{"a", "b"}
	expect(s).To.Have.Len(2)
	expect(s).Not.To.Have.Len(1)
	c := make(chan bool, 5)
	c <- true
	expect(c).To.Have.Len(1)
	expect(c).Not.To.Have.Len(0)
}

func TestCap(t *testing.T) {
	expect := expect.New(t)
	expect([2]int{}).To.Have.Cap(2)
	expect(make([]byte, 2, 10)).To.Have.Cap(10)
	expect(make(chan string, 2)).Not.To.Have.Cap(10)
}

func TestKey(t *testing.T) {
	expect := expect.New(t)
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	expect(m1).To.Have.Key("a")
	expect(m1).Not.To.Have.Key("c")
	expect(m1).To.Have.Key("a", 1)

	m2 := map[int]string{
		1: "a",
		2: "b",
	}
	expect(m2).To.Have.Key(1)
	expect(m2).Not.To.Have.Key(3)
	expect(m2).To.Have.Key(2, "b")
	expect(m2).Not.To.Have.Key(1, "c")

	m3 := map[string]interface{}{
		"arr": [1]int{},
		"map": map[int]int{1: 1},
	}
	expect(m3).To.Have.Key("arr")
	expect(m3).To.Have.Key("map")
	expect(m3).Not.To.Have.Key("struct")
	expect(m3).To.Have.Key("arr", [1]int{})
	expect(m3).To.Have.Key("map", map[int]int{1: 1})
	expect(m3).Not.To.Have.Key("map", map[string]int{})
}

func TestKeys(t *testing.T) {
	expect := expect.New(t)
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	expect(m1).To.Have.Keys("a", "b", "c")
	expect(m1).Not.To.Have.Keys("d", "e", "i")

	m2 := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	expect(m2).To.Have.Keys(1, 2, 3)
	expect(m2).Not.To.Have.Keys(4, 5, 6)
}

func TestField(t *testing.T) {
	expect := expect.New(t)
	p := struct {
		X, Y int
	}{1, 3}
	expect(p).To.Have.Field("X")
	expect(p).To.Have.Field("Y", 3)
	expect(p).Not.To.Have.Field("Z")
	expect(p).Not.To.Have.Field("Y", 4)
}

func TestFields(t *testing.T) {
	expect := expect.New(t)
	p := struct {
		X, Y int
	}{1, 2}
	expect(p).To.Have.Fields("X", "Y")
	expect(p).Not.To.Have.Fields("Z")
	expect(p).Not.To.Have.Fields("T", "Z")
}

// Test Method
type Person struct{}

func (p Person) Hello()  {}
func (p *Person) Hallo() {}

func TestMethod(t *testing.T) {
	expect := expect.New(t)
	p := Person{}
	expect(p).To.Have.Method("Hello")
	expect(p).Not.To.Have.Method("Hallo")
	expect(&p).To.Have.Method("Hallo")
	expect(&p).To.Have.Method("Hello")
}

func TestHaveFailNow(t *testing.T) {
	mockT := newMockT()
	expect := expect.New(mockT)
	l := []string{"foo"}
	expect(l).To.Have.Len(1).Else.FailNow()
	select {
	case <-mockT.FailNowCalled:
		t.Fatalf("Expected FailNow() on passing test not to be called")
	default:
	}
	expect(l).To.Have.Len(3).Else.FailNow()
	select {
	case <-mockT.FailNowCalled:
	default:
		t.Fatalf("Expected FailNow() on failing test to be called")
	}
}

func TestNotHaveFailNow(t *testing.T) {
	mockT := newMockT()
	expect := expect.New(mockT)
	l := []string{"foo"}
	expect(l).Not.To.Have.Len(3).Else.FailNow()
	select {
	case <-mockT.FailNowCalled:
		t.Fatalf("Expected FailNow() on passing test not to be called")
	default:
	}
	expect(l).Not.To.Have.Len(1).Else.FailNow()
	select {
	case <-mockT.FailNowCalled:
	default:
		t.Fatalf("Expected FailNow() on failing test to be called")
	}
}
