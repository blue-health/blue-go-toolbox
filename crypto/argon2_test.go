package crypto

import (
	"bytes"
	"encoding/base64"
	"testing"
)

func TestArgonHash(t *testing.T) {
	testCases := []struct{ in string }{
		{in: ""},
		{in: "helloworld"},
		{in: "..123,,,12mk1p."},
	}

	for _, c := range testCases {
		hash, err := ArgonHash(c.in, DefaultArgonParams)
		if err != nil {
			t.Fatal("ArgonHash returned error", err)
		}

		ok, err := ArgonCompareHash(c.in, hash)
		if err != nil {
			t.Fatal("ArgonCompareHash returned error", err)
		}

		if !ok {
			t.Fatal("ArgonCompareHash returned false.")
		}
	}
}

func TestArgonDecodeHash(t *testing.T) {
	type decodeOut struct {
		params ArgonParams
		salt   string
		hash   string
		err    error
	}

	testCases := []struct {
		in  string
		out decodeOut
	}{
		{
			in: "$$$$",
			out: decodeOut{
				ArgonParams{}, "", "", ErrInvalidHash,
			},
		},
		{
			in: "$argon2id$v=500000$m=65536,t=3,p=2$Woo1mErn1s7AHf96ewQ8Uw$D4TzIwGO4XD2buk96qAP+Ed2baMo/KbTRMqXX00wtsU",
			out: decodeOut{
				ArgonParams{}, "", "", ErrIncompatibleVersion,
			},
		},
		{
			in: dummyHash,
			out: decodeOut{
				DefaultArgonParams, "Woo1mErn1s7AHf96ewQ8Uw", "D4TzIwGO4XD2buk96qAP+Ed2baMo/KbTRMqXX00wtsU", nil,
			},
		},
	}

	for _, c := range testCases {
		params, salt, hash, err := decodeHash(c.in)
		if err != nil {
			if c.out.err == err {
				continue
			}

			t.Fatal("decodeHash returned unexpected error", err)
		}

		if params != c.out.params {
			t.Fatal("decodeHash returned false params. Expected:", c.out.params, "Got:", params)
		}

		expectSalt, err := base64.RawStdEncoding.Strict().DecodeString(c.out.salt)
		if err != nil {
			panic(err)
		}

		expectHash, err := base64.RawStdEncoding.Strict().DecodeString(c.out.hash)
		if err != nil {
			panic(err)
		}

		if !bytes.Equal(salt, expectSalt) {
			t.Fatal("decodeHash returned false salt. Expected:", expectSalt, "Got:", salt)
		}

		if !bytes.Equal(hash, expectHash) {
			t.Fatal("decodeHash returned false hash. Expected:", expectHash, "Got:", hash)
		}
	}
}
