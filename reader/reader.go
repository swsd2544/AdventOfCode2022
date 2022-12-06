package reader

import (
	"os"
)

func GetTextFromInputFile(filename string) (string, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
