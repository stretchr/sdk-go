package stretchr

func (s *Session) LoadOne(path string) (*Resource, error) {

	res, err := s.session.At(path).Read()

	if err != nil {
		return nil, err
	}

	return nil, nil
}
