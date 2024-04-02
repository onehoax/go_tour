package myinterface

import (
	"fmt"
	"io"
	"strings"
)

func Reader() {
	var r *strings.Reader = strings.NewReader("This message goes into a buffer 8 bytes at a time!")

	var b []byte = make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}

type myReader struct{}

// reader that emits an infinite stream of the ASCII char 'A'// reader that emits an infinite stream of the ASCII char 'A“
func (r *myReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 'A' // 65
	}
	return len(b), nil
}

func Readcustom() {
	r := &myReader{}
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}
