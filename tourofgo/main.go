package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) (n int, err error) {
	cipher := make([]byte, 1024)
	read := 0
	for {
		n, err := rot13.r.Read(cipher)
		for c := 0; c < n; c++ {
			if cipher[c] >= 65 && cipher[c] <= 90 {
				// preserve case. this is upper
				b[read+c] = 65 + (cipher[c]-65+13)%26
			} else if cipher[c] >= 97 && cipher[c] <= 122 {
				// preserve case. this is lower
				b[read+c] = 97 + (cipher[c]-97+13)%26
			} else {
				b[read+c] = cipher[c]
			}
		}
		read += n
		if err == io.EOF {
			break
		}
	}
	return read, io.EOF
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
