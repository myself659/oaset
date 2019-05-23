package main

import "fmt"

func main() {
	s1 := "4579K"
	s2 := "AAAA2"
	s1ok := checkstr(s1)
	if !s1ok {
		fmt.Println("invalid string")
		return
	}
	s2ok := checkstr(s2)
	if !s2ok {
		fmt.Println("invalid string")
		return
	}
	fmt.Println(compare(s1, s2))
}

var validchars = map[byte]bool{
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
	'A': true,
	'J': true,
	'Q': true,
	'K': true,
}

func checkstr(s string) bool {
	if len(s) != 5 {
		return false
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		ok := validchars[c]
		if !ok {
			return false
		}
	}
	return true
}

func getvalue(c byte) int {
	switch c {
	case '2', '3', '4', '5', '6', '7', '8', '9':
		{
			return int(c - '0')
		}
	case 'A':
		{
			return 11
		}
	case 'J', 'Q', 'K':
		{
			return 10
		}
	}

	return -1
}

func calniu(nums []int) int {
	var ret = -11

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := 0
			niucnt := 0
			maxcard := 2
			for k := 0; k < len(nums); k++ {
				if k == i {
					niucnt += nums[k]
					continue
				}
				if k == j {
					niucnt += nums[k]
					continue
				}
				if nums[k] > 10 {
					sum += (nums[k] - 10)
				} else {
					sum += nums[k]
				}
				if nums[k] > maxcard {
					maxcard = nums[k]
				}

			}
			if sum%10 == 0 {
				niucnt = niucnt % 10
				if niucnt == 0 {
					niucnt = 10
				}
			} else {
				niucnt = maxcard - 12
			}
			if niucnt > ret {
				ret = niucnt
			}
		}
	}
	return ret
}

func getnums(s string) []int {
	var ret []int

	for i := 0; i < len(s); i++ {
		value := getvalue(byte(s[i]))

		ret = append(ret, value)
	}

	return ret
}

func compare(s1 string, s2 string) int {
	num1 := getnums(s1)
	num2 := getnums(s2)
	fmt.Println("num1:", num1)
	fmt.Println("num2:", num2)
	niu1 := calniu(num1)
	niu2 := calniu(num2)
	fmt.Println("niu1:", niu1)
	fmt.Println("niu2:", niu2)
	if niu1 == niu2 {
		return 0
	}

	if niu1 > niu2 {
		return 1
	}

	return -1
}
