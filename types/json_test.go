package types_test

import (
	"bytes"
	"testing"

	"github.com/blue-health/blue-go-toolbox/types"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

type testStruct struct {
	JS types.JSON `yaml:"yamlJSON"`
}

func TestJSONMarshalYAML(t *testing.T) {
	testCases := []struct {
		name string
		in   types.JSON
		out  string
		err  error
	}{
		{
			name: "normal json",
			in: types.JSON{
				JSON:  []byte(`{"hello": "world"}`),
				Valid: true,
			},
			out: `yamlJSON: '{"hello":"world"}'
`,
		},
		{
			name: "null json",
			in: types.JSON{
				JSON:  []byte(`{"hello": "world"}`),
				Valid: false,
			},
			out: `yamlJSON: "null"
`,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			ts := &testStruct{JS: c.in}

			out, err := yaml.Marshal(&ts)
			if err != nil || c.err != nil {
				if c.err != nil {
					require.ErrorIs(t, err, c.err)
				} else {
					require.NotNil(t, err)
				}

				return
			}

			require.Equal(t, c.out, string(out))
		})
	}
}

func TestJSONUnmarshalYAML(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  types.JSON
		err  error
	}{
		{
			name: "normal json",
			in:   `yamlJSON: '{"hello": "world"}'`,
			out: types.JSON{
				JSON:  []byte(`{"hello": "world"}`),
				Valid: true,
			},
		},
		{
			name: "null json",
			in:   `yamlJSON: ''`,
			out: types.JSON{
				JSON:  []byte(`{"hello": "world"}`),
				Valid: false,
			},
		},
		{
			name: "multiline json",
			in: `yamlJSON: | 
                {"hello": "world"}`,
			out: types.JSON{
				JSON:  []byte(`{"hello": "world"}`),
				Valid: false,
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			var ts testStruct

			err := yaml.Unmarshal([]byte(c.in), &ts)
			require.Nil(t, err)

			if c.out.Valid {
				require.Equal(t, c.out, ts.JS)
			}
		})
	}
}

func TestRemoveJSONKey(t *testing.T) {
	cases := []struct {
		old types.JSON
		new types.JSON
		key string
	}{
		{
			old: types.JSON{},
			new: types.JSON{},
			key: "none",
		},
		{
			old: types.JSONFrom([]byte(`{"some":"object"}`)),
			new: types.JSONFrom([]byte(`{"some":"object"}`)),
			key: "none",
		},
		{
			old: types.JSONFrom([]byte(`{"some":"object"}`)),
			new: types.JSONFrom([]byte(`{}`)),
			key: "some",
		},
		{
			old: types.JSONFrom([]byte(`{"some":{"nested":"object"}}`)),
			new: types.JSONFrom([]byte(`{"some":{}}`)),
			key: "nested",
		},
	}

	for i := range cases {
		n := types.RemoveJSONKey(cases[i].old, cases[i].key)

		if !bytes.Equal(cases[i].new.JSON, n.JSON) {
			t.Fatalf("JSON objects %s and %s don't match", string(cases[i].old.JSON), string(n.JSON))
		}
	}
}
