package types

import (
	"errors"

	"github.com/jackc/pgtype"
)

type JSON pgtype.JSON

var ErrJSONInvalid = errors.New("json_invalid")

func (j JSON) MarshalYAML() ([]byte, error) {
	return j.Bytes, nil
}

func (j *JSON) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	if j == nil {
		return ErrJSONInvalid
	}

	j.Bytes = append(j.Bytes, s...)

	return nil
}
