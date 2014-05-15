package common

/*
	HTTP Methods
*/

const (
	// HTTPMethodGet represents the HTTP Method for GET requests, literally "GET".
	HTTPMethodGet string = "GET"
	// HTTPMethodPost represents the HTTP Method for POST requests, literally "POST".
	HTTPMethodPost string = "POST"
	// HTTPMethodPut represents the HTTP Method for PUT requests, literally "PUT".
	HTTPMethodPut string = "PUT"
	// HTTPMethodPatch represents the HTTP Method for PATCH requests, literally "PATCH".
	HTTPMethodPatch string = "PATCH"
	// HTTPMethodDelete represents the HTTP Method for DELETE requests, literally "DELETE".
	HTTPMethodDelete string = "DELETE"
)

/*
	URL and Path segments
*/

const (
	// FilterFieldPrefix is the character that identifies a filter field.
	FilterFieldPrefix string = ":"

	// ProtocolSeparator is the string used to separate the protocol from the rest of the URL.
	ProtocolSeparator string = "://"

	// PathSeparator is the character that separates path segments.
	PathSeparator string = "/"

	// HostSeparator is the character that separates the account from the host.
	HostSeparator string = "."

	// TopLevelHostName is the default host name.
	TopLevelHostName string = "stretchr.com"

	// APIVersionPathPrefix is the string defining the api and version prefix in the URL
	APIVersionPathPrefix string = "/api/v"
)

/*
	Standard response object fields
*/

// The keys used in the standard respnose object
const (
	ResponseObjectFieldData          string = "~data"
	ResponseObjectFieldChangeInfo    string = "~changes"
	ResponseObjectFieldErrors        string = "~errors"
	ResponseObjectFieldStatusCode    string = "~status"
	ResponseObjectFieldContext       string = "~context"
	ResponseObjectFieldErrorsMessage string = "~message"
	ResponseObjectFieldTotal         string = "~total"
)

// The keys used in the change info object
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
	// DataFieldID is the ~id field key
	DataFieldID string = "~id"
)

/*
	Modifiers
*/

// The modifiers available to manipulate the request.
const (
	ModifierLimit string = "limit"
	ModifierSkip  string = "skip"
	ModifierOrder string = "order"
)

/*
	Signing
*/
const (
	// ParameterAPIKey is the key string for constructing the URL
	ParameterAPIKey string = "key"
)

/*
	HTTP Protocols
*/

// The available HTTP protocols
const (
	HTTPProtocol       string = "http"
	HTTPProtocolSecure string = "https"
)
