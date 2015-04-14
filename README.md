# Expect.go
> Minimalistic BDD-style assertions for Go (inspired by expect.js)

```go
expect := expect.New(t)

// Numbers
expect(10).To.Be.Above(1).And.Below(20)
expect(5).Not.To.Be.Within(0, 4)

// Strings
expect("foobarbaz").To.StartWith("foo").And.EndWith("baz").And.Contains("bar")
expect("Foo").To.Match("(?i)foo")

// Equal
expect(false).Not.To.Equal("false")
expect(map[int]int{}).To.Equal(map[int]int{})

// Empty
expect(map[int]int{}).To.Be.Empty()
expect("").To.Be.Empty()
expect([2]int{}).Not.To.Be.Empty()

// Ok (i.e: not "", 0, false, nil)
expect(val).To.Be.Ok()
expect(false).Not.To.Be.Ok()

// Type Assertion
expect("").To.Be.String()
expect(0).To.Be.Int()
expect(1.1).To.Be.Float()
expect(1).Not.To.Be.Bool()
expect(map[string]int{}).To.Be.Map()
expect([...]int{1}).To.Be.Array()
expect([]string{"a"}).To.Be.Slice()
expect(ch).To.Be.Chan()
expect(struct{}{}).To.Be.Struct()
expect(&struct{}{}).To.Be.Ptr()
expect(nil).To.Be.Nil()
expect(Person{}).To.Be.Type("Person")

// Len
expect("foo").To.Have.Len(3)
expect([]int{1, 2}).To.Have.Len(2)

// Cap
expect(make([]byte, 2, 10)).To.Have.Cap(10)
expect([2]int{}).To.Have.Cap(2)

// Maps
m1 := map[string]int{
	"a": 1,
	"b": 2,
}
expect(m1).To.Have.Key("a")
expect(m1).To.Have.Key("a", 1) // With value
expect(m1).To.Have.Keys("a", "b")

// Structs
p := struct {
	X, Y int
}{1, 3}
expect(p).To.Have.Field("Y", 3).And.Field("X", 1)
```

## License
MIT
