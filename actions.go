package stretchr

func (s *Session) LoadOne(path string) (*Resource, error) {

	res, err := s.session.At(path).Read()

	// TODO: Carry on

	return nil, nil
}
