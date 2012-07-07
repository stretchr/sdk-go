package stretchr

type Resource map[string]interface{}

func MakeResource() Resource {
	return make(Resource)
}
