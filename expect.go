package expect

// T is a type that we can perform assertions with.
type T interface {
	Errorf(format string, args ...interface{})
	Fatal(...interface{})
	FailNow()
}

// Matcher is any type which can perform matches against an actual
// value.  A non-nil error means a failed match, and its Error()
// method should return a suffix that finishes the prefix,
// "Expected {actual} to ".
//
// Note that for negated matches (expect(foo).Not...), the matcher
// itself will be printed out using %#v syntax.  If you would like
// to customize your output, implement fmt.GoStringer.
//
// Example:
//
//     type StatusCodeMatcher int
//
//     func (m StatusCodeMatcher) Match(actual interface{}) error {
//         resp, ok := actual.(*http.Response)
//         if !ok {
//             // "Expected {actual} to be of type *http.Response"
//             return errors.New("be of type *http.Response")
//         }
//         if resp.StatusCode != int(m) {
//             // "Expected {actual} to have a response code {m}"
//             return fmt.Errorf("have a response code of %d", m)
//         }
//         return nil
//     }
//
//     // GoString returns representation for m in a negated state,
//     // e.g. where a value was *not* supposed to match.
//     func (m StatusCodeMatcher) GoString() string {
//         return fmt.Sprintf("Status Code %d", m)
//     }
//
//     // The eventual assertions:
//     expect(resp).To.Pass(StatusCodeMatcher(http.StatusOK))
//     expect(resp).Not.To.Pass(StatusCodeMatcher(http.StatusOK))
type Matcher interface {
	Match(actual interface{}) error
}

// Expect is the result of calling an Expectation with a value to assert
// against.
type Expect struct {
	To  *To
	Not *Not
}

// Not stores a To that is initialized to be the negation of the main To.
type Not struct {
	To *To
}

// Expectation returns an Expect that is used to make assertions.
type Expectation func(v interface{}) *Expect

// New returns a function that can be given a value and have `To, To.Be,
// To.Have` assertions chained onto it.
func New(t T) Expectation {
	return func(v interface{}) *Expect {
		return &Expect{
			To:  newTo(t, v, true),
			Not: &Not{To: newTo(t, v, false)},
		}
	}
}
