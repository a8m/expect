// +build !go1.7

package expect

import (
	"fmt"
	"runtime"
)

// stackPrefix returns all stack lines.  This is the go1.6
// and earlier version, which is reportedly less accurate
// than the logic from 1.7+.
func stackPrefix(stack []uintptr) []string {
	var output []string
	for _, pc := range stack {
		file, line := runtime.FuncForPC(pc).FileLine()
		output = append(output, fmt.Sprintf("%s:%d", file, line))
	}
	return output
}
