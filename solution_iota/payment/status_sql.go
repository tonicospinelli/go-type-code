package payment

import (
	"database/sql/driver"
)

func (s *Status) Scan(value interface{}) error {
	switch value := value.(type) {
	case nil:
		return nil
	case string:
		return s.UnmarshalText([]byte(value))
	case int:
		*s = Status(value)
		return nil
	}

	return errStatusNotAvailable{status: value}

}

func (s *Status) Value() (driver.Value, error) {
	if s == nil {
		return nil, nil
	}
	return s.String(), nil
}
