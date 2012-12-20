package stretchr

import (
	"bytes"
)

func MergeStrings(stringArray ...string) string {

	var buffer bytes.Buffer
	for _, v := range stringArray {
		buffer.WriteString(v)
	}
	return buffer.String()

}

func JoinStrings(separator string, stringArray ...string) string {

	var buffer bytes.Buffer
	var max int = len(stringArray) - 1
	for vi, v := range stringArray {
		buffer.WriteString(v)
		if vi < max {
			buffer.WriteString(separator)
		}
	}
	return buffer.String()

}
