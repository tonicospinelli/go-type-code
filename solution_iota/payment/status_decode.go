package payment

func (s *Status) UnmarshalBinary(data []byte) error {
	return s.UnmarshalText(data)
}

func (s *Status) UnmarshalJSON(data []byte) error {
	return s.UnmarshalText(data)
}

func (s *Status) UnmarshalText(data []byte) error {
	p, err := Parse(string(data))

	*s = p
	return err
}
