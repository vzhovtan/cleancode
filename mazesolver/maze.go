package main

import "fmt"

type Point struct {
	x int
	y int
}

var direction = [][]int{
	{1, 0},  //top
	{0, 1},  //right
	{0, -1}, //bottom
	{-1, 0}, //left
}

func walk(maze [][]string, wall string, curr Point, exit Point, seen [][]bool, path *[]Point) bool {
	//1.location is off the map
	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
		return false
	}

	//2.meet a wall
	if maze[curr.y][curr.x] == wall {
		return false
	}

	//3.reached exit
	if curr.x == exit.x && curr.y == exit.y {
		*path = append(*path, exit)
		return true
	}

	//4.reached visited point
	if seen[curr.x][curr.y] {
		return false
	}

	//recursion
	seen[curr.y][curr.x] = true
	*path = append(*path, curr)

	for i := 0; i < len(direction); i++ {
		a, b := direction[i][0], direction[i][1]
		newCurr := Point{
			x: curr.x + a,
			y: curr.y + b,
		}
		if walk(maze, wall, newCurr, exit, seen, path) {
			return true
		}
	}
	item := len(*path) - 1
	rm(path, item)

	return false
}

func rm(slc *[]Point, item int) {
	s := *slc
	s = s[:item]
	*slc = s
}

func main() {
	maze := [][]string{
		{"########## #"},
		{"########## #"},
		{"#########  #"},
		{"#######  ###"},
		{"####### ####"},
	}
	wall := "#"
	start := Point{
		x: 4,
		y: 7,
	}
	exit := Point{
		x: 0,
		y: 10,
	}
	seen := [][]bool{
		{false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false},
	}
	path := []Point{}

	walk(maze, wall, start, exit, seen, &path)
	fmt.Printf("%v\n", path)
}
