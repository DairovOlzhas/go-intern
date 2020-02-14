package main

import "fmt"

func main(){
	//arr := [6]int{1,2,3,4,5,6}
	//slice := arr[0:3]
	//slice = append(slice, 5,5,6,6,6)
	//fmt.Println(slice)
	//fmt.Println("length: ", len(slice))
	//fmt.Println("capacity: ", cap(slice))

	//slice = append(slice, 5,5,6,6,6)
	//fmt.Println(slice)
	//fmt.Println("length: ", len(slice))
	//fmt.Println("capacity: ", cap(slice))

	//slice = append([]int{1,2,3}, slice...)
	//fmt.Println(slice)

	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%v is %v\n", i, IsOdd(i))
	//}

	//d := DoubleX()
	//fmt.Println(d())
	//fmt.Println(d())


	//dino := &Dino(){}
	//dino.Walk()
	//PrintWalking(dino)

	nl := make(chan int)

	go printNumbers(nl,0)
	go printNumbers(nl,10)

	for i := 0; i < 20; i++ {
		fmt.Println(<-nl)
	}
}

func IsOdd(n int) string{
	if n%2 == 1 {
		return "odd"
	} else {
		return "even"
	}
}

func swap(a,b int) (c,d int) {
	c = b
	d = a
	return
}

func add(a *[]int, b ...int) {
	*a = append(*a, b...)
}

func DoubleX() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func Fibo(x int) int {
	if x < 2 {
		return 1
	}
	return Fibo(x - 1) + Fibo(x - 2)
}

type DinoInt interface {
	Walk()
}

type Dino struct {
	Name string
}

func (d *Dino) Walk() {
	fmt.Println("Walking")
}

func PrintWalking(d DinoInt){
	d.Walk()
}


func printNumbers(a chan <- int, n int) {
	for i := n; i < n+10; i++ {
		a <- i
	}
}