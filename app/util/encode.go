package util

import (
	"encoding/base64"
	"io"
	"os"
)

func EncodeFile(file *os.File) (string, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	encodedContent := base64.StdEncoding.EncodeToString(content)
	return encodedContent, nil
}
