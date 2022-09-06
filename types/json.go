package types

import (
	"errors"

	"github.com/jackc/pgtype"
)

type JSON pgtype.JSON

var ErrJSONInvalid = errors.New("json_invalid")

func JSONFromBytes(b []byte) JSON {
	return JSON{
		Bytes:  b,
		Status: pgtype.Present,
	}
}

func (j JSON) ToBytes() []byte {
	if j.Status == pgtype.Present {
		return j.Bytes
	}

	return []byte{}
}

func (j JSON) Valid() bool {
	return j.Status == pgtype.Present
}

func (j JSON) MarshalYAML() (interface{}, error) {
	switch j.Status {
	case pgtype.Present:
		return string(j.Bytes), nil
	case pgtype.Null:
		return "", nil
	case pgtype.Undefined:
		return nil, ErrJSONInvalid
	}

	return nil, ErrJSONInvalid
}

func (j *JSON) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	if s == "" || s == "null" {
		*j = JSON{
			Status: pgtype.Null,
		}
	} else {
		*j = JSON{
			Bytes:  []byte(s),
			Status: pgtype.Present,
		}
	}

	return nil
}
