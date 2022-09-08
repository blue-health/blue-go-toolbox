package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"github.com/bojanz/currency"
)

type Currency struct{ currency.Amount }

var ErrCurrencyInvalid = errors.New("currency_invalid")

func Rounded(c currency.Amount) Currency {
	return Currency{Amount: c.RoundTo(2, currency.RoundHalfUp)}
}

func Abs(c currency.Amount) (Currency, error) {
	if c.IsNegative() {
		a, err := c.Mul("-1")
		if err != nil {
			return Currency{}, fmt.Errorf("failed to calculate absolute value: %w", err)
		}

		return Currency{Amount: a}, nil
	}

	return Currency{Amount: c}, nil
}

func (c Currency) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%s)", c.Number(), c.CurrencyCode()), nil
}

func (c *Currency) Scan(src interface{}) error {
	var source string

	switch v := src.(type) {
	case string:
		source = v
	case []byte:
		source = string(v)
	default:
		return ErrCurrencyInvalid
	}

	if source == "" {
		return ErrCurrencyInvalid
	}

	values := strings.Split(strings.Trim(source, "()"), ",")

	if len(values) != 2 {
		return ErrCurrencyInvalid
	}

	var (
		amount = values[0]
		code   = values[1]
	)

	n, err := currency.NewAmount(amount, code)
	if err != nil {
		return err
	}

	c.Amount = n

	return nil
}

func (c Currency) MarshalYAML() (interface{}, error) {
	return struct {
		Number       string `yaml:"number"`
		CurrencyCode string `yaml:"currency"`
	}{
		Number:       c.Amount.Number(),
		CurrencyCode: c.Amount.CurrencyCode(),
	}, nil
}

func (c *Currency) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var crc struct {
		Number       string `yaml:"number"`
		CurrencyCode string `yaml:"currency"`
	}

	if err := unmarshal(&crc); err != nil {
		return err
	}

	n, err := currency.NewAmount(crc.Number, crc.CurrencyCode)
	if err != nil {
		return err
	}

	c.Amount = n

	return nil
}

func MustParseCurrency(value, currencyCode string) currency.Amount {
	n, err := currency.NewAmount(value, currencyCode)
	if err != nil {
		panic(err)
	}

	return n
}
