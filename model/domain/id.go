package domain

import (
	"database/sql/driver"
	"errors"

	uuid "github.com/satori/go.uuid"
)

// ID of domain implement for uuid
type ID [16]byte

// NewID create new UUID with v4
func NewID() ID {
	return ID(uuid.NewV4())
}

// IDFromString convert string to ID
func IDFromString(s string) (ID, error) {
	id, err := uuid.FromString(s)
	return ID(id), err
}

// IsZero check uuid is zero
func (u *ID) IsZero() bool {
	if u == nil {
		return true
	}
	for _, c := range u {
		if c != 0 {
			return false
		}
	}
	return true
}

func (u ID) String() string {
	return uuid.UUID(u).String()
}

// MarshalJSON implement for json encoding
func (u ID) MarshalJSON() ([]byte, error) {
	if len(u) == 0 {
		return []byte(`""`), nil
	}
	return []byte(`"` + u.String() + `"`), nil
}

// UnmarshalJSON implement for json decoding
func (u *ID) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == `""` {
		return nil
	}

	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("invalid UUID format")
	}
	data = data[1 : len(data)-1]
	uu, err := uuid.FromString(string(data))
	if err != nil {
		return errors.New("invalid UUID format")
	}
	*u = ID(uu)
	return nil
}

// Scan .
func (u *ID) Scan(b interface{}) error {
	if b == nil {
		for i := range u {
			u[i] = 0
		}
		return nil
	}

	// postgres store DB as a string
	id, err := uuid.FromString(string(b.([]byte)))
	if err != nil {
		return err
	}

	for i, c := range id {
		u[i] = c
	}

	return nil
}

// Value .
func (u ID) Value() (driver.Value, error) {
	if u.IsZero() {
		return nil, nil
	}
	return uuid.UUID(u).String(), nil
}
