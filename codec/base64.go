package codec

import (
	"encoding/base64"
)

func ToBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func FromBase64(s string) string {
	bytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return s
	}
	return string(bytes)
}
