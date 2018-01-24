package demos

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		fmt.Printf("b[%v]: %v\n", i, b[i])
		b[i] += 13
	}
	return n, err
}

// RotReaderTest function
func RotReaderTest() {
	s := strings.NewReader("Hello World")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
