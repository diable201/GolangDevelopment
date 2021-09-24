package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (n int, e error) {
	n, e = rot13.r.Read(b)
	for i := range b {
		if (b[i] >= 'a' && b[i] < 'n') || (b[i] >= 'A' && b[i] < 'N') {
			b[i] += 13
		} else if (b[i] >= 'n' && b[i] <= 'z') || (b[i] >= 'N' && b[i] <= 'Z') {
			b[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
