package payment

func (s Status) MarshalBinary() ([]byte, error) {
	return s.MarshalText()
}

func (s Status) MarshalJSON() ([]byte, error) {
	return s.MarshalText()
}

func (s Status) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}
