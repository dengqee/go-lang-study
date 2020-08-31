package main

import (
	"fmt"
	"io"
	"os"
)

type CountWriter struct {
	w io.Writer
	n int64
}

func (c *CountWriter) Write(p []byte) (int, error) {
	c.n = int64(len(p))
	ret, err := c.w.Write(p)
	return ret, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	ret := &CountWriter{w: w}
	return ret, &ret.n

}

func main() {
	cw, l := CountingWriter(os.Stdout)
	cw.Write([]byte("hello"))
	fmt.Println(*l)
}
