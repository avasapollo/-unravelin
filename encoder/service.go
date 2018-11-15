package encoder

import (
	"crypto/sha1"
	"encoding/base64"
	"hash"
)

type hashEncoder struct {
	encoder hash.Hash
}

func NewHashEncoder() HashEncoder {
	return &hashEncoder{
		encoder: sha1.New(),
	}
}

func (h hashEncoder) Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func (h hashEncoder) Decode(s string) (string, error) {
	res, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(res[:]), nil
}
