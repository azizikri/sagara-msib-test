package models

import (
	"database/sql/driver"
	"errors"
)

// Size enum
type Size string

// Enum values for Size
const (
	Small  Size = "S"
	Medium Size = "M"
	Large  Size = "L"
)

// Validate the size to ensure it's valid
func (s Size) IsValid() error {
	switch s {
	case Small, Medium, Large:
		return nil
	}
	return errors.New("invalid Size")
}

func (s *Size) Scan(value interface{}) error {
	strValue, ok := value.(string)
	if !ok {
		return errors.New("invalid data for Size")
	}
	*s = Size(strValue)
	return nil
}

func (s Size) Value() (driver.Value, error) {
	if err := s.IsValid(); err != nil {
		return nil, err
	}
	return string(s), nil
}
