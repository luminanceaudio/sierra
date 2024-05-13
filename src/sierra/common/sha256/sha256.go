package sha256

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

const (
	sha256Length = 64
)

type Sha256 = string

// New parses a sha256 string into Sha256 type
func New(sha256 string) (Sha256, error) {
	if len(sha256) != sha256Length {
		return "", fmt.Errorf("sha256 length must be %d", sha256Length)
	}

	return sha256, nil
}

func HashHexEncoded(reader io.ReadSeeker) (*Sha256, error) {
	_, err := reader.Seek(0, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("failed to seek to start while hashing")
	}

	hasher := sha256.New()
	_, err = io.Copy(hasher, reader)
	if err != nil {
		return nil, fmt.Errorf("failed hashing file: %w", err)
	}
	sha256Str := hex.EncodeToString(hasher.Sum(nil))
	return &sha256Str, nil
}
