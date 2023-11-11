package readers

import (
	"bufio"
	"io"
)

func ReadUrls(buf io.Reader, ch chan string) error {
	scanner := bufio.NewScanner(buf)

	for scanner.Scan() {
		ch <- scanner.Text()
	}

	return nil
}
