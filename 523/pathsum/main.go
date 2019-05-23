package main

import (
	"fmt"
	"math"
)

func main() {

	row1 := []int{3, 0, -2, 4, 0}
	row2 := []int{-1, 2, -2, 1, 4}
	row3 := []int{3, 1, -2, -3, 3}
	row4 := []int{2, -4, -3, -3, 2}
	row5 := []int{5, 2, -2, -3, 1}
	grid := [][]int{}
	grid = append(grid, row1)
	grid = append(grid, row2)
	grid = append(grid, row3)
	grid = append(grid, row4)
	grid = append(grid, row5)

	res := pathsum(grid)
	fmt.Println("res:", res)
	return
}

func pathsum(grid [][]int) int {
	ymax := len(grid)
	if ymax == 0 {
		return 0
	}
	xmax := len(grid[0])
	dp := make([][]int, ymax)
	for i := range dp {
		dp[i] = make([]int, xmax)
	}

	visited := make([][]bool, ymax)
	for i := range dp {
		visited[i] = make([]bool, xmax)
	}
	for x := 0; x < xmax; x++ {
		dp[0][x] = grid[0][x]
		visited[0][x] = true
	}
	for y := 1; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			xpos := x - 2
			ypos := y - 1
			tempsum := math.MaxInt32
			cnt := 0
			if xpos >= 0 && xpos < xmax && ypos >= 0 && ypos < ymax && visited[ypos][xpos] == true {
				sum := dp[ypos][xpos] + grid[y][x]
				if sum < tempsum {
					tempsum = sum
				}
				cnt++
			}
			xpos = x + 2
			ypos = y - 1
			if xpos >= 0 && xpos < xmax && ypos >= 0 && ypos < ymax && visited[ypos][xpos] == true {
				sum := dp[ypos][xpos] + grid[y][x]
				if sum < tempsum {
					tempsum = sum
				}
				cnt++
			}

			xpos = x - 1
			ypos = y - 2
			if xpos >= 0 && xpos < xmax && ypos >= 0 && ypos < ymax && visited[ypos][xpos] == true {
				sum := dp[ypos][xpos] + grid[y][x]
				if sum < tempsum {
					tempsum = sum
				}
				cnt++
			}

			xpos = x + 1
			ypos = y - 2
			if xpos >= 0 && xpos < xmax && ypos >= 0 && ypos < ymax && visited[ypos][xpos] == true {
				sum := dp[ypos][xpos] + grid[y][x]
				if sum < tempsum {
					tempsum = sum
				}
				cnt++
			}
			// 支持加跳的话，增加新的方向即可，即对应增加y+1, y+2的方向
			if cnt > 0 {
				visited[y][x] = true
				dp[y][x] = tempsum
			}
		}
	}
	fmt.Println("DP:", dp)
	fmt.Println("visited:", visited)
	ret := math.MaxInt32
	for x := 0; x < xmax; x++ {
		if visited[ymax-1][x] {
			if ret > dp[ymax-1][x] {
				ret = dp[ymax-1][x]
			}
		}
	}
	return ret
}
