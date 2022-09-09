package types

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type UUIDArray []uuid.UUID

func (a UUIDArray) Value() (driver.Value, error) {
	s := make(pq.StringArray, len(a))

	for i := range a {
		s[i] = a[i].String()
	}

	return s.Value()
}

func (a *UUIDArray) Scan(src interface{}) error {
	var s pq.StringArray
	if err := s.Scan(src); err != nil {
		return err
	}

	*a = make(UUIDArray, len(s))

	for i := range s {
		u, err := uuid.Parse(s[i])
		if err != nil {
			return err
		}

		(*a)[i] = u
	}

	return nil
}
