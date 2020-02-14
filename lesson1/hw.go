package main

import "fmt"


func main(){
	a:="a33bc1//4/g6a/510d"
	var last rune
	var num int
	for i := 0; i < len(a); i++ {
		if a[i] >= '0' && a[i] <= '9' {
			num*=10
			num+=int(a[i]-'0')
		} else {
			for num > 1 {
				fmt.Printf("%v", string(last))
				num--
			}
			num = 0
			if a[i] == '/' && i+1 < len(a){
				i++
				last = rune(a[i])
				fmt.Printf("%v", string(last))
			} else {
				last = rune(a[i])
				fmt.Printf("%v", string(last))
			} 
		}
 	}
}

