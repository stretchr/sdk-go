package stretchr

/*
	HTTP Methods
*/

const (
	// HttpMethodGet represents the HTTP Method for GET requests, literally "GET".
	HttpMethodGet string = "GET"
	// HttpMethodGet represents the HTTP Method for POST requests, literally "POST".
	HttpMethodPost string = "POST"
	// HttpMethodGet represents the HTTP Method for PUT requests, literally "PUT".
	HttpMethodPut string = "PUT"
	// HttpMethodGet represents the HTTP Method for DELETE requests, literally "DELETE".
	HttpMethodDelete string = "DELETE"
)

/*
	URL and Path segments
*/

const (
	filterFieldPrefix string = ":"

	protocolSeparator string = "://"

	pathSeparator string = "/"

	hostSeparator string = "."

	topLevelHostName string = "stretchr.com"

	apiVersionPathPrefix string = "/api/v"
)

/*
	Data fields
*/

const (
	dataFieldID string = "~id"
)

/*
	Modifiers
*/

const (
	modifierLimit string = "~limit"
)

/*
	HTTP Protocols
*/
const (
	httpProtocol       string = "http"
	httpProtocolSecure string = "https"
)
