package expect

import "testing"

// TODO(Ariel): Create mock that implement TB interface
// and stub `Error` and `Fatal`

func TestLen(t *testing.T) {
	expect := New(t)
	expect("foo").To.Have.Len(3)
	m := map[string]int{}
	expect(m).To.Have.Len(0)
	expect(m).Not.To.Have.Len(1)
}

func TestKey(t *testing.T) {
	expect := New(t)
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
