package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println(isNumber("7e1000"))

}

func isNumber(s string) bool {
	var ss string
	l,r := 0, len(s)-1
	for l < len(s) && s[l] == ' ' {l++}
	for r >= 0 && s[r] == ' ' {r--}
	if r < l {
		return false
	}

	for i:=l; i <= r; i++ {
		ss += string(s[i])
	}

	_, err := strconv.ParseFloat(ss, 1024)

	if err != nil {
		return false
	}
	return true

}


