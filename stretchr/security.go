package stretchr

import (
	"crypto/sha1"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

const (
	// FailedSignature is the string that is returned when signing fails.
	FailedSignature string = "(failed signature)"
)

// Hash generates an SHA1 hash of the specified string.
func Hash(s string) string {

	hash := sha1.New()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", hash.Sum(nil))

}

// getOrderedParams gets the parameters ordered by key, then by values.
func getOrderedParams(values url.Values) string {

	// get the keys
	var keys []string
	for k, _ := range values {
		keys = append(keys, k)
	}

	// sort the keys
	sort.Strings(keys)

	// ordered items
	var ordered []string

	// sort the values
	for _, key := range keys {
		sort.Strings(values[key])
		for _, val := range values[key] {
			ordered = append(ordered, fmt.Sprintf("%s=%s", url.QueryEscape(key), url.QueryEscape(val)))
		}
	}

	joined := strings.Join(ordered, "&")
	return joined

}

// GetSignature gets the signature of a request based on the given parameters.
func GetSignature(method, requestUrl, body, privateKey string) (string, error) {

	// parse the URL
	u, parseErr := url.ParseRequestURI(requestUrl)

	if parseErr != nil {
		return FailedSignature, parseErr
	}

	// get the query values
	values := u.Query()

	// add the private key parameter
	values.Set(PrivateKeyKey, privateKey)

	if len(body) > 0 {
		values.Set(BodyHashKey, Hash(body))
	}

	// get the ordered params
	orderedParams := getOrderedParams(values)

	base := strings.Split(u.String(), "?")[0]
	combined := fmt.Sprintf("%s&%s?%s", strings.ToUpper(method), base, orderedParams)

	return Hash(combined), nil

}

// GetSignedURL gets the URL with the sign parameter added based on the given parameters.
func GetSignedURL(method, requestUrl, body, privateKey string) (string, error) {

	hash, hashErr := GetSignature(method, requestUrl, body, privateKey)

	if hashErr != nil {
		return FailedSignature, hashErr
	}

	signed := fmt.Sprintf("%s&%s=%s", requestUrl, url.QueryEscape(SignatureKey), url.QueryEscape(hash))

	return signed, nil

}
