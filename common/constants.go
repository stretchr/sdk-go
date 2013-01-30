package common

/*
	HTTP Methods
*/

const (
	// common.HttpMethodGet represents the HTTP Method for GET requests, literally "GET".
	HttpMethodGet string = "GET"
	// common.HttpMethodGet represents the HTTP Method for POST requests, literally "POST".
	HttpMethodPost string = "POST"
	// common.HttpMethodGet represents the HTTP Method for PUT requests, literally "PUT".
	HttpMethodPut string = "PUT"
	// common.HttpMethodGet represents the HTTP Method for DELETE requests, literally "DELETE".
	HttpMethodDelete string = "DELETE"
)

/*
	URL and Path segments
*/

const (
	FilterFieldPrefix string = ":"

	ProtocolSeparator string = "://"

	PathSeparator string = "/"

	HostSeparator string = "."

	TopLevelHostName string = "stretchr.com"

	ApiVersionPathPrefix string = "/api/v"
)

/*
	Data fields
*/

const (
	DataFieldID string = "~id"
)

/*
	Modifiers
*/

const (
	ModifierLimit string = "~limit"
	ModifierSkip  string = "~skip"
	ModifierOrder string = "~order"
)

/*
	Signing
*/

const (
	SignSignature  string = "~sign"
	SignPrivateKey string = "~private"
	SignPublicKey  string = "~key"
	SignBodyHash   string = "~bodyhash"
)

/*
	HTTP Protocols
*/
const (
	HttpProtocol       string = "http"
	HttpProtocolSecure string = "https"
)

/*
	Headers
*/
const (
	HeaderResponseHash string = "X-Stretchr-Response-Hash"
)
