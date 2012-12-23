package api

import (
	"bytes"
	"fmt"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestStringy_MergeStrings(t *testing.T) {

	assert.Equal(t, "callback(jsonString)", MergeStrings("callback", "(", "jsonString", ")"))

}

func TestStringy_MergeBytes(t *testing.T) {

	assert.Equal(t, []byte("callback(jsonString)"), MergeBytes([]byte("callback"), []byte("("), []byte("jsonString"), []byte(")")))

}

func TestStringy_Join(t *testing.T) {

	assert.Equal(t, "projects/centivus/accounts/tyler", JoinStrings("/", "projects", "centivus", "accounts", "tyler"))

}

func Benchmark_SprintF(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("projects/%s/accounts/%s", string(i), string(i))
	}

}

func Benchmark_Join(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = JoinStrings("/", "projects", string(i), "accounts", string(i))
	}

}

func Benchmark_Bytes(b *testing.B) {

	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		buffer.WriteString("/")
		buffer.WriteString("projects")
		buffer.WriteString(string(i))
		buffer.WriteString("accounts")
		buffer.WriteString(string(i))
		buffer.String()
	}

}
