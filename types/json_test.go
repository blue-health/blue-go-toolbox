package types_test

import (
	"testing"

	"github.com/blue-health/blue-go-toolbox/types"
	"github.com/jackc/pgtype"
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
				Bytes:  []byte(`{"hello": "world}`),
				Status: pgtype.Present,
			},
			out: `yamlJSON: '{"hello": "world}'
`,
		},
		{
			name: "null json",
			in: types.JSON{
				Bytes:  []byte(`{"hello": "world}`),
				Status: pgtype.Null,
			},
			out: `yamlJSON: ""
`,
		},
		{
			name: "undefined json",
			in: types.JSON{
				Bytes:  []byte(`{"hello": "world}`),
				Status: pgtype.Undefined,
			},
			out: `yamlJSON: '{"hello": "world}'`,
			err: types.ErrJSONInvalid,
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
			in:   `yamlJSON: '{"hello": "world}'`,
			out: types.JSON{
				Bytes:  []byte(`{"hello": "world}`),
				Status: pgtype.Present,
			},
		},
		{
			name: "null json",
			in:   `yamlJSON: ''`,
			out: types.JSON{
				Bytes:  []byte(`{"hello": "world}`),
				Status: pgtype.Null,
			},
		},
		{
			name: "multiline json",
			in: `yamlJSON: | 
                {"hello": "world"}`,
			out: types.JSON{
				Bytes:  []byte(`{"hello": "world"}`),
				Status: pgtype.Present,
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			var ts testStruct

			err := yaml.Unmarshal([]byte(c.in), &ts)
			require.Nil(t, err)

			if c.out.Status == pgtype.Present {
				require.Equal(t, c.out, ts.JS)
			}
		})
	}
}
