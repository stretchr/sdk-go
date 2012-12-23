package stretchr

// Resource describes objects that can be used in conjunction with Stretchr services.
type Resource interface {

	// ResourcePath gets the path for this Resource.
	ResourcePath() string

	// ResourceData gets the data for this Resource.
	ResourceData() map[string]interface{}
}
