package crypto

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type ArgonParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

const (
	dummyText = "immunkarte-for-the-win-12345"
	dummyHash = "$argon2id$v=19$m=65536,t=3,p=2$Woo1mErn1s7AHf96ewQ8Uw$D4TzIwGO4XD2buk96qAP+Ed2baMo/KbTRMqXX00wtsU"
)

var (
	ErrInvalidHash         = errors.New("invalid_argon_hash")
	ErrIncompatibleVersion = errors.New("incompatible_argon_version")
	DefaultArgonParams     = ArgonParams{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
)

func ArgonHash(password string, p ArgonParams) (string, error) {
	n, err := RandomBytes(uint(p.saltLength))
	if err != nil {
		return "", fmt.Errorf("failed to get random bytes: %w", err)
	}

	var (
		h       = argon2.IDKey([]byte(password), n, p.iterations, p.memory, p.parallelism, p.keyLength)
		b64Salt = base64.RawStdEncoding.EncodeToString(n)
		b64Hash = base64.RawStdEncoding.EncodeToString(h)
	)

	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash), nil
}

func ArgonCompareHash(password, encodedPassword string) (bool, error) {
	p, salt, hash, err := decodeHash(encodedPassword)
	if err != nil {
		return false, fmt.Errorf("failed to decode argon2 hash: %w", err)
	}

	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	return subtle.ConstantTimeCompare(hash, otherHash) == 1, nil
}

func ArgonDummyCompare() {
	_, _ = ArgonCompareHash(dummyText, dummyHash)
}

func decodeHash(encodedHash string) (params ArgonParams, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return ArgonParams{}, nil, nil, ErrInvalidHash
	}

	var version int
	if _, err = fmt.Sscanf(vals[2], "v=%d", &version); err != nil {
		return ArgonParams{}, nil, nil, fmt.Errorf("failed to scan argon2 version: %w", err)
	}

	if version != argon2.Version {
		return ArgonParams{}, nil, nil, ErrIncompatibleVersion
	}

	var p ArgonParams
	if _, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism); err != nil {
		return ArgonParams{}, nil, nil, fmt.Errorf("failed to scan argon2 params: %w", err)
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return ArgonParams{}, nil, nil, fmt.Errorf("failed to decode argon2 salt: %w", err)
	}

	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return ArgonParams{}, nil, nil, fmt.Errorf("failed to decode argon2 hash: %w", err)
	}

	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}
