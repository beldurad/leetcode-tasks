package main

import "fmt"

type Point struct {
	i int
	j int
}

func numIslands(grid [][]byte) int {
	islandMap := make(map[Point]int) // Map, which returns island num by given point
	islandCount := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				continue
			}
			curPoint := Point{
				i: i,
				j: j,
			}
			if _, ok := islandMap[curPoint]; ok {
				continue
			}
			islandCount++
			neighborsDetection(i, j, islandCount, grid, islandMap)
		}
	}
	return islandCount
}

func neighborsDetection(i, j, island int, grid [][]byte, islandMap map[Point]int) {
	if !isInRange(i, j, grid) || grid[i][j] == '0' {
		return
	}
	curPoint := Point{
		i: i,
		j: j,
	}
	if _, ok := islandMap[curPoint]; ok {
		return
	}
	islandMap[curPoint] = island
	neighbors := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	for _, n := range neighbors {
		neighborsDetection(i+n[0], j+n[1], island, grid, islandMap)
	}
}

func isInRange(i, j int, grid [][]byte) bool {
	if i < 0 || i >= len(grid) {
		return false
	}
	if j < 0 || j >= len(grid[i]) {
		return false
	}
	return true
}

func main() {
	fmt.Println(numIslands([][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}}))
}
