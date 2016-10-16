// +build !go1.7

package expect

// T is a type that we can perform assertions with.  This is the
// pre-go1.7 version, which does not have a Run function.
type T interface {
	Errorf(format string, args ...interface{})
	Fatal(...interface{})
	FailNow()
}
