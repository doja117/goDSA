package main

import "fmt"

func solve(s, p string) {
	if len(s) == 0 {
		fmt.Printf("%s\t", p)
		return
	}

	solve(s[1:], p)
	p += string(s[0])
	solve(s[1:len(s)], p)

}
func main() {
	solve("abc", "")
}
