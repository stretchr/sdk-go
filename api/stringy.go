package api

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

func MergeBytes(byteArray ...[]byte) []byte {

	var buffer bytes.Buffer
	for _, v := range byteArray {
		buffer.Write(v)
	}
	return buffer.Bytes()

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
