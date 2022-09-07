package types

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSON struct {
	JSON  json.RawMessage
	Valid bool
}

var (
	nullBytes      = []byte("null")
	ErrJSONInvalid = errors.New("json_invalid")
)

func JSONFromBytes(j []byte) JSON {
	return JSON{JSON: j, Valid: true}
}

func (j JSON) Value() (driver.Value, error) {
	if !j.Valid {
		return nil, nil
	}

	bs, err := json.Marshal(j.JSON)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (j *JSON) Scan(src interface{}) error {
	var source []byte

	switch t := src.(type) {
	case string:
		if t == "" {
			source = nullBytes
		} else {
			source = []byte(t)
		}
	case []byte:
		if len(t) == 0 {
			source = nullBytes
		} else {
			source = t
		}
	case nil:
		source = nullBytes
	default:
		return ErrJSONInvalid
	}

	if bytes.Equal(source, nullBytes) {
		*j = JSON{JSON: nullBytes, Valid: false}
	} else {
		var bs json.RawMessage
		if err := json.Unmarshal(source, &bs); err != nil {
			return err
		}

		*j = JSON{JSON: bs, Valid: true}
	}

	return nil
}

func (j JSON) MarshalJSON() ([]byte, error) {
	if !j.Valid {
		return nullBytes, nil
	}

	bs, err := json.Marshal(j.JSON)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (j *JSON) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, nullBytes) {
		*j = JSON{JSON: nullBytes, Valid: false}
	} else {
		var bs json.RawMessage
		if err := json.Unmarshal(data, &bs); err != nil {
			return err
		}

		*j = JSON{JSON: bs, Valid: true}
	}

	return nil
}

func (j *JSON) Unmarshal(v interface{}) error {
	return j.Scan(v)
}

func (j JSON) MarshalYAML() (interface{}, error) {
	if !j.Valid {
		return string(nullBytes), nil
	}

	bs, err := json.Marshal(j.JSON)
	if err != nil {
		return nil, err
	}

	return string(bs), nil
}

func (j *JSON) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	source := []byte(s)

	if len(source) == 0 || bytes.Equal(source, nullBytes) {
		*j = JSON{JSON: nullBytes, Valid: false}
	} else {
		var bs json.RawMessage
		if err := json.Unmarshal(source, &bs); err != nil {
			return err
		}

		*j = JSON{JSON: bs, Valid: true}
	}

	return nil
}
