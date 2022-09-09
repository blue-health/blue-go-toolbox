package types_test

import (
	"testing"

	"github.com/blue-health/blue-go-toolbox/types"
	"github.com/bojanz/currency"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

type testCurrencyStruct struct {
	C types.Currency `yaml:"yamlCurrency"`
}

func TestCurrencyMarshalYAML(t *testing.T) {
	a, err := currency.NewAmount("12.90", "EUR")
	require.NoError(t, err)

	b, err := currency.NewAmount("22.90", "EUR")
	require.NoError(t, err)

	testCases := []struct {
		name string
		in   types.Currency
		out  string
	}{
		{
			name: "normal amount 1",
			in:   types.CurrencyFrom(a),
			out:  "yamlCurrency:\n    number: \"12.90\"\n    currency: EUR\n",
		},
		{
			name: "normal amount 2",
			in:   types.CurrencyFrom(b),
			out:  "yamlCurrency:\n    number: \"22.90\"\n    currency: EUR\n",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			ts := &testCurrencyStruct{C: c.in}

			out, err := yaml.Marshal(&ts)
			require.NoError(t, err)
			require.Equal(t, string(out), c.out)
		})
	}
}
