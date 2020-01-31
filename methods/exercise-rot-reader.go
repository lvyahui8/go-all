package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13r *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13r.r.Read(b)
	if err != nil {
		//if err == io.EOF {
		//	return 0, nil
		//}
		return 0, err
	}
	for i := 0; i < n; i++ {
		b[i] = b[i] + 1
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
