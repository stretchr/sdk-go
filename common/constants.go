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
	Standard response object fields
*/

const (
	ResponseObjectFieldData          string = "~data"
	ResponseObjectFieldChangeInfo    string = "~changes"
	ResponseObjectFieldErrors        string = "~errors"
	ResponseObjectFieldStatusCode    string = "~status"
	ResponseObjectFieldContext       string = "~context"
	ResponseObjectFieldErrorsMessage string = "~message"
	ItemsFieldKey                    string = "~items"
)

const (
	ChangeInfoPublicFieldCreated      string = "~created"
	ChangeInfoPublicFieldUpdated      string = "~updated"
	ChangeInfoPublicFieldDeleted      string = "~deleted"
	ChangeInfoPublicFieldDeltas       string = "~deltas"
	ChangeInfoPublicFieldDeltaCreated string = "~created"
	ChangeInfoPublicFieldDeltaUpdated string = "~updated"
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
	ModifierLimit string = "limit"
	ModifierSkip  string = "skip"
	ModifierOrder string = "order"
)

/*
	Signing
*/

const (
	SignSignature  string = "sign"
	SignPrivateKey string = "private"
	SignAPIKey  string = "key"
	SignBodyHash   string = "bodyhash"
)

/*
	HTTP Protocols
*/
const (
	HttpProtocol       string = "http"
	HttpProtocolSecure string = "https"
)
