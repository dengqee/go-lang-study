package intset

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint32
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/32, uint(x%32)
	return word < len(s.words) && s.words[word]&1<<bit != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/32, uint(x%32)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Len() int {
	return len(s.words) * 32
}

func (s *IntSet) Remove(x int) {
	word, bit := x/32, uint(x%32)
	s.words[word] &= ^(1 << bit)
	if s.words[0] == 0 {
		if len(s.words) > 1 {
			s.words = s.words[1:]
		} else {
			s.words = []uint32{}
		}
	}
}

func (s *IntSet) Clear() {
	s.words = []uint32{}
}

func (s *IntSet) Copt() *IntSet {
	ret := *s

	return &ret
}
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 32; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 32*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
func (s *IntSet) AddAll(input ...int) {
	for _, i := range input {
		s.Add(i)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}
}
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}

}

func (s *IntSet) Elems() []int {
	var ret []int
	for i := len(s.words) - 1; i >= 0; i-- {
		word := s.words[i]
		for b := 0; b < 32; b++ {
			if word&1<<b != 0 {
				ret = append(ret, (len(s.words)-1-i)*32+b)
			}
		}
	}
	return ret
}
