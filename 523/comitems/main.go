package main

import "fmt"

func main() {
	num1 := []int{1, 3, 5, 7, 9}
	num2 := []int{3, 5, 7, 9, 11}
	res := comitems(num1, num2)
	fmt.Println("res:", res)
}

func comitems(num1 []int, num2 []int) []int {
	var ret []int
	num1len := len(num1)
	num2len := len(num2)
	i := 0
	j := 0
	for i < num1len && j < num2len {
		if num1[i] == num2[j] {
			ret = append(ret, num1[i])
			i++
			j++
			continue
		}
		if num1[i] > num2[j] {
			j++
			continue
		}
		i++
	}

	return ret
}
