package types

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"strings"

	"github.com/google/uuid"
)

type UUID struct {
	UUID  uuid.UUID
	Valid bool
}

func FromUUID(id uuid.UUID) UUID {
	return UUID{UUID: id, Valid: true}
}

func (u UUID) Value() (driver.Value, error) {
	if !u.Valid {
		return nil, nil
	}

	return u.UUID.Value()
}

func (u *UUID) Scan(src interface{}) error {
	u.UUID, u.Valid = uuid.Nil, false

	switch t := src.(type) {
	case string:
		if t == "" {
			return nil
		}
	case []byte:
		if len(t) == 0 {
			return nil
		}
	case nil:
		return nil
	}

	u.Valid = true

	return u.UUID.Scan(src)
}

func (u UUID) MarshalJSON() ([]byte, error) {
	if u.Valid {
		return json.Marshal(u.UUID.String())
	}

	return nullBytes, nil
}

func (u *UUID) UnmarshalJSON(text []byte) error {
	u.UUID, u.Valid = uuid.Nil, false

	if bytes.Equal(text, nullBytes) {
		return nil
	}

	p, err := uuid.Parse(strings.Trim(string(text), "\""))
	if err != nil {
		return err
	}

	u.UUID, u.Valid = p, true

	return nil
}

func (u *UUID) UnmarshalText(text []byte) error {
	return u.UnmarshalJSON(text)
}

func (u *UUID) UnmarshalYAML(unmarshal func(interface{}) error) error {
	u.UUID, u.Valid = uuid.Nil, false

	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	p, err := uuid.Parse(s)
	if err != nil {
		return err
	}

	u.UUID, u.Valid = p, true

	return nil
}
