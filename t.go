// +build go1.7

package expect

import "testing"

// T is a type that we can perform assertions with.
type T interface {
	Run(name string, test func(t *testing.T)) (succeeded bool)
	Errorf(format string, args ...interface{})
	Fatal(...interface{})
	FailNow()
}
