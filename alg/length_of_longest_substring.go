package main

import (
	"fmt"
	"strings"
)

type Empty struct {
}

var empty Empty

type Set struct {
	m map[rune]Empty
}

func SetFactory() *Set {
	return &Set{
		m: map[rune]Empty{},
	}
}

func (s *Set) Add(val rune) {
	s.m[val] = empty
}
func (s *Set) Remove(val rune) {
	delete(s.m, val)
}
func (s *Set) Len() int {
	return len(s.m)

}
func (s *Set) Clear() {
	s.m = make(map[rune]Empty)

}
func (s *Set) Traverse() {
	for v := range s.m {
		fmt.Println(v)
	}

}
func (s *Set) Contains(val rune) bool {
	for v := range s.m {
		if v == val {
			return true
		}
	}
	return false
}
func main() {

}
func lenOfLongestSubstring(s string) int {
	var res int
	var curStr string
	left := 0
	right := 0
	curStr = s[left:right]
	for ; right < len(s); right++ {
		if index := strings.IndexByte(curStr, s[right]); index != -1 {
			left += index + 1
		}
		curStr = s[left : right+1]
		if len(curStr) > res {
			res = len(curStr)
		}
	}
	return res
}
