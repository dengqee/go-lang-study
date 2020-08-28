package intset
import (
	"fmt"
	"bytes"
)
type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.word[word]&1<<bit != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.word, 0)
	}
	s.words[word] |= 1 << bit
}
func (s *IntSet) String() string {
    var buf bytes.Buffer
    buf.WriteByte('{')
    for i, word := range s.words {
        if word == 0 {
            continue
        }
        for j := 0; j < 64; j++ {
            if word&(1<<uint(j)) != 0 {
                if buf.Len() > len("{") {
                    buf.WriteByte(' ')
                }
                fmt.Fprintf(&buf, "%d", 64*i+j)
            }
        }
    }
    buf.WriteByte('}')
    return buf.String()
}
func (s *IntSet) Len() int {
	return len(*s) * 64
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &= ^(1 << bit)
	if s.words[0] == 0 {
		if len(s.words) > 1 {
			s.words = s.words[1:]
		} else {
			s.words = make([]uint64)
		}
	}
}

func (s *IntSet) Clear() {
	s.words = make([]uint64)
}

func (s *IntSet) Copt() *IntSet {
	ret := *s

	return &ret
}

func (s *IntSet) AddAll(input ...int) {
	for _, i := range input {
		s.Add(i)
	}
}
