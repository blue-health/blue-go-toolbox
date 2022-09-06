package crypto

import (
	"bytes"
	"testing"
)

func TestAESEncryptDecrypt(t *testing.T) {
	key, err := RandomBytes(32)
	if err != nil {
		panic(err)
	}

	a, err := NewAES(key)
	if err != nil {
		panic(err)
	}

	testCases := []struct{ in []byte }{
		{in: nil},
		{in: []byte("")},
		{in: []byte("helloworld")},
		{in: []byte("...")},
	}

	for _, c := range testCases {
		en, err := a.Encrypt(c.in)
		if err != nil {
			t.Fatal("AES.Encrypt returned error", err)
		}

		de, err := a.Decrypt(en)
		if err != nil {
			t.Fatal("AES.Decrypt returned error", err)
		}

		if !bytes.Equal(c.in, de) {
			t.Fatal("Encrypt/Decrypt are mismatched. Expected:", c.in, "Got:", de)
		}
	}
}
