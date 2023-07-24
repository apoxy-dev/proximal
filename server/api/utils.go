package api

import (
	"encoding/base64"
	"fmt"
)

func decodeNextPageToken(token string) (string, error) {
	dTs := make([]byte, base64.StdEncoding.DecodedLen(len(token)))
	n, err := base64.StdEncoding.Decode(dTs, []byte(token))
	if err != nil {
		return "", fmt.Errorf("failed to decode page token: %v", err)

	}
	dTs = dTs[:n]
	return string(dTs), nil
}

func encodeNextPageToken(ts string) string {
	return base64.StdEncoding.EncodeToString([]byte(ts))
}
