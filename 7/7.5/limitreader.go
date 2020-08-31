package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

type LimiteReader struct {
	r io.Reader
	n int64
}

func (r *LimiteReader) Read(p []byte) (n int, err error) {
	if r.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > r.n {
		p = p[:r.n]
	}
	n, err = r.r.Read(p)
	r.n -= int64(n)
	return
}
func LimitRead(r io.Reader, n int64) io.Reader {
	return &LimiteReader{r, n}
}

func main() {
	s := []byte("aaabbdddc")
	r := bytes.NewReader(s)
	reader := LimitRead(r, 3)
	s1, _ := ioutil.ReadAll(reader)
	fmt.Println(string(s1))
}
