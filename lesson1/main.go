package main

import "fmt"


// const a float32 = 123

func main(){

	// var a, b int = 3, 4
	// fmt.Println(a,b)
	// fmt.Printf("a = %v, b = %v \n", a, b)
	// var b int = 5

	

	// var c string
	// var d float32
	
	// a := 5 //динамичесикй
	// fmt.Println(a.type)
	// fmt.Println("c is = " + c)
	// fmt.Println("d is = %x", d)



	// var a rune = 48
	// fmt.Println(string(a))


	// var a string = "hello" // значение стринга статично
	// d := "hello world"

	// fmt.Println(d)


	// a := 8969
	// fmt.Println(string(a))

	// a := 'a'
	// a = 'd'
	// fmt.Println("%s", a)

	// if a > 6 {
	// 	fmt.Println("Less")
	// } else {
	// 	fmt.Println("OK")
	// }


	// for i:=0; i < 10; i++ {

	// }


	// b := "hello world"

	// for i, v := range b {
	// 	fmt.Printf("Index is : %v \n", i)
	// 	fmt.Printf("Values is: %v \n", string(v))
	// }


	// b := "hello world"
	// for _, v := range b {
	// 	fmt.Printf("Values is: %v \n", string(v))
	// }



	// a := [10]int{}
	// b := make([].int, 5)

	// for i:=0; i < 10; i++ {
	// 	a[i] = i
	// }

	// b := a[1:4]
	// b[2] = 9
	// fmt.Println(b)
	// fmt.Println(a)

	// c := 6
	// d := &c

	// fmt.Println(c)

	// *d = 16

	// fmt.Println(d)

	// fmt.Println(c)



	// c := "hello"


	// var a int64 = 100000000000000000000000
	// fmt.Println(a)


	// HOMework
	// a:="a3bc4g6a\\6"

	a:="a3bc4g6a/510d"
	// "aaabccccgggggga\\\\\\\"
	// fmt.Println(a)

	// for _, v := range a {
	// 	fmt.Printf("%v \n", v)
	// }

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

			if a[i] == '/' {
				if i+1 < len(a) {
					i++
					last = rune(a[i])
					fmt.Printf("%v", string(last))
				} 
			} else {
				last = rune(a[i])
				fmt.Printf("%v", string(last))
			} 
		}
 	}

	// m, err = strconv.Atoi("5")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(m)

	// var l interface{} = 5
	// var l interface{} = 5.5

	// switch l.(type) {
	// case int:
	// 	fmt.Println("Values is integer")
	// case float32:
	// 	fmt.Println("Values is float")
	// default:
	// 	fmt.Println("Values is %v", l.(type))
	// }

}