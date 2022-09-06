package types

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type UUID pgtype.UUID

func FromUUID(u uuid.UUID) UUID {
	return UUID{
		Bytes:  u,
		Status: pgtype.Present,
	}
}

func (u UUID) ToUUID() uuid.UUID {
	if u.Status == pgtype.Present {
		return u.Bytes
	}

	return uuid.Nil
}
