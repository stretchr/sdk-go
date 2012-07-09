package stretchr

// Delete deletes this resource.
//
// For example:
//  deleteErr := resource.Delete()
func (r *Resource) Delete() error {
	return r.session.Delete(r.path, r.GetID())
}
