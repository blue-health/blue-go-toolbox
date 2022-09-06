package crypto

import (
	"testing"
)

func TestHMACGenerateValidate(t *testing.T) {
	key, err := RandomBytes(32)
	if err != nil {
		panic(err)
	}

	h := NewHMAC(key)

	testCases := []struct{ in []byte }{
		{in: nil}, {in: []byte("")},
		{in: []byte("helloworld")},
		{in: []byte("...")},
	}

	for _, c := range testCases {
		en := h.Generate(c.in)

		ok := h.Validate(c.in, en)
		if !ok {
			t.Fatal("Generate/Validate are mismatched.")
		}
	}
}
