package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type TaxRate struct {
	Rate uint8  `json:"rate" yaml:"rate"`
	Kind string `json:"kind" yaml:"kind"`
}

var ErrTaxRateInvalid = errors.New("tax_rate_invalid")

func (t TaxRate) Value() (driver.Value, error) {
	return fmt.Sprintf("(%d,%s)", t.Rate, t.Kind), nil
}

func (t *TaxRate) Scan(src interface{}) error {
	var source string

	switch v := src.(type) {
	case string:
		source = v
	case []byte:
		source = string(v)
	default:
		return ErrTaxRateInvalid
	}

	if source == "" {
		return ErrTaxRateInvalid
	}

	values := strings.Split(strings.Trim(source, "()"), ",")

	if len(values) != 2 {
		return ErrTaxRateInvalid
	}

	var (
		rate = values[0]
		kind = values[1]
	)

	i, err := strconv.Atoi(rate)
	if err != nil {
		return ErrTaxRateInvalid
	}

	if !(i >= 0 && i < 100) {
		return ErrTaxRateInvalid
	}

	t.Rate = uint8(i)
	t.Kind = kind

	return nil
}
