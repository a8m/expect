// +build go1.7

package expect

import (
	"fmt"
	"runtime"
)

// stackPrefix returns all stack lines.
func stackPrefix(stack []uintptr) []string {
	var output []string
	frames := runtime.CallersFrames(stack)
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		output = append(output, fmt.Sprintf("%s:%d", frame.File, frame.Line))
	}
	return output
}
