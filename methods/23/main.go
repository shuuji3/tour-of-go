package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (int, error) {
	tmp := make([]byte, len(b))
	if _, err := reader.r.Read(tmp); err != nil {
		return 0, err
	}

	for i, char := range tmp {
		if 'a' <= char && char <= 'z' {
			b[i] = (char-'a'+13)%26 + 'a'
		} else if 'A' <= char && char <= 'Z' {
			b[i] = (char-'A'+13)%26 + 'A'
		} else {
			b[i] = char
		}
	}
	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
