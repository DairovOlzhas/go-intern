package main

import "fmt"

func main(){
	fmt.Println(reverse(123))
}

func twoSum(nums []int, target int) []int {
	res := make([]int,0)
	for i:= 0; i < len(nums); i++ {
		for j:=i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				res = append(res, i,j)
				return res
			}
		}
	}
	return res
}

func reverse(x int) int {
	if x < -2147483648 && x > 2147483647 {
		return 0
	}
	var res int
	for x > 0 {

		res *= 10
		res += x%10
		x /= 10

		//fmt.Println(res)
	}
	return res
}