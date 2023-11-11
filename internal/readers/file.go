package readers

import (
	"os"
)

func ReadFile(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return f, nil
}
