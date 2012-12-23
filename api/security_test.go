package stretchr

import (
	"github.com/stretchrcom/testify/assert"
	"net/url"
	"testing"
)

func TestHash(t *testing.T) {

	bytes := []byte("abc")

	assert.Equal(t, "a9993e364706816aba3e25717850c26c9cd0d89d", hash(bytes), "MAIN Hash test")

}

func TestGetOrderedParams(t *testing.T) {

	values := make(url.Values)
	values.Add("~key", "ABC123")
	values.Add(":name", "!Mat")
	values.Add(":name", "!Laurie")
	values.Add(":age", ">20")
	values.Add(":something", ">2 0")

	ordered := getOrderedParams(values)

	assert.Equal(t, "%3Aage=%3E20&%3Aname=%21Laurie&%3Aname=%21Mat&%3Asomething=%3E2+0&~key=ABC123", ordered)

}

func TestgetSignature(t *testing.T) {

	var signed string

	signed, _ = getSignature(HttpMethodGet, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", []byte("body"), "ABC123-private")
	assert.Equal(t, "df073ee4086eed5848d167871c7424937027728e", signed)

	signed, _ = getSignature("get", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", []byte("body"), "ABC123-private")
	assert.Equal(t, "df073ee4086eed5848d167871c7424937027728e", signed, "Lower case method shouldn't affect GetSignature")

	signed, _ = getSignature(HttpMethodGet, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", []byte("body"), "DIFFERENT-PRIVATE")
	assert.Equal(t, "34f55c3a086c260098e75066b38ac42e33e8faab", signed)

	signed, _ = getSignature(HttpMethodGet, "http://test.stretchr.com/api/v1?:name=!Laurie&~key=ABC123&:age=>20&:name=!Mat", []byte("body"), "DIFFERENT-PRIVATE")
	assert.Equal(t, "34f55c3a086c260098e75066b38ac42e33e8faab", signed, "Different order of args shouldn't matter")

}

func TestGetSignedURL(t *testing.T) {

	var signed string

	signed, _ = getSignedURL(HttpMethodGet, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", []byte("body"), "ABC123-private")
	assert.Equal(t, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=df073ee4086eed5848d167871c7424937027728e", signed)

	signed, _ = getSignedURL(HttpMethodGet, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", []byte("body"), "DIFFERENT-PRIVATE")
	assert.Equal(t, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=34f55c3a086c260098e75066b38ac42e33e8faab", signed)

}

func TestNoBodyHashWhenNoBody(t *testing.T) {

	signed, _ := getSignedURL(HttpMethodGet, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", []byte(""), "ABC123-private")
	assert.Equal(t, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=bdf49047abf3c8e56de21e244bc24b1c2a6086a2", signed)

}
