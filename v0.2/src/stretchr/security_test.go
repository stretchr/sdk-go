package stretchr

import (
	"net/url"
	"testing"
)

func TestHash(t *testing.T) {

	str := "abc"

	AssertEqual(t, "a9993e364706816aba3e25717850c26c9cd0d89d", Hash(str), "MAIN Hash test")

}

func TestGetOrderedParams(t *testing.T) {

	values := make(url.Values)
	values.Add("~key", "ABC123")
	values.Add(":name", "!Mat")
	values.Add(":name", "!Laurie")
	values.Add(":age", ">20")
	values.Add(":something", ">2 0")

	ordered := getOrderedParams(values)

	AssertEqual(t, "%3Aage=%3E20&%3Aname=%21Laurie&%3Aname=%21Mat&%3Asomething=%3E2+0&~key=ABC123", ordered)

}

func TestGetSignature(t *testing.T) {

	var signed string

	signed, _ = GetSignature(GetMethod, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "ABC123-private")
	AssertEqual(t, "df073ee4086eed5848d167871c7424937027728e", signed)

	signed, _ = GetSignature("get", "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "ABC123-private")
	AssertEqual(t, "df073ee4086eed5848d167871c7424937027728e", signed, "Lower case method shouldn't affect GetSignature")

	signed, _ = GetSignature(GetMethod, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "DIFFERENT-PRIVATE")
	AssertEqual(t, "34f55c3a086c260098e75066b38ac42e33e8faab", signed)

	signed, _ = GetSignature(GetMethod, "http://test.stretchr.com/api/v1?:name=!Laurie&~key=ABC123&:age=>20&:name=!Mat", "body", "DIFFERENT-PRIVATE")
	AssertEqual(t, "34f55c3a086c260098e75066b38ac42e33e8faab", signed, "Different order of args shouldn't matter")

}

func TestGetSignedURL(t *testing.T) {

	var signed string

	signed, _ = GetSignedURL(GetMethod, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "ABC123-private")
	AssertEqual(t, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=df073ee4086eed5848d167871c7424937027728e", signed)

	signed, _ = GetSignedURL(GetMethod, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20", "body", "DIFFERENT-PRIVATE")
	AssertEqual(t, "http://test.stretchr.com/api/v1?~key=ABC123&:name=!Mat&:name=!Laurie&:age=>20&~sign=34f55c3a086c260098e75066b38ac42e33e8faab", signed)

}
