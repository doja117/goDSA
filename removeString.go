package main

import "fmt"

func stringSolve(s, c string) string {
	if len(s) < len(c) {
		return s
	}
	t := s[:len(c)]
	//fmt.Println(t, s)
	if t == c {
		return stringSolve(s[len(c):], c)
	}
	return string(s[0]) + stringSolve(s[1:], c)
}

func main() {
	s1 := "Saurabh Ojha Kumar"
	s2 := "Kumar"
	fmt.Println(stringSolve(s1, s2))
}
