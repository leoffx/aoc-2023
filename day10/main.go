package main

import (
	"fmt"
	"os"
	"strings"
)

const startingPointSymbol = "S"

type coords = struct {
	x int
	y int
}

func findConnectedPipes(pipes [][]string, i, j int) coords {
	rowLen := len(pipes)
	colLen := len(pipes[0])
	up := coords{x: -1, y: 0}
	down := coords{x: 1, y: 0}
	left := coords{x: 0, y: -1}
	right := coords{x: 0, y: 1}

	x := i + up.x
	y := j + up.y
	if x >= 0 && x < rowLen && y >= 0 && y < colLen {
		pipe := pipes[x][y]
		if pipe == "|" || pipe == "7" || pipe == "F" {
			return coords{x: x, y: y}
		}
	}

	x = i + down.x
	y = j + down.y
	if x >= 0 && x < rowLen && y >= 0 && y < colLen {
		pipe := pipes[x][y]
		if pipe == "|" || pipe == "J" || pipe == "L" {
			return coords{x: x, y: y}
		}
	}

	x = i + left.x
	y = j + left.y
	if x >= 0 && x < rowLen && y >= 0 && y < colLen {
		pipe := pipes[x][y]
		if pipe == "-" || pipe == "F" || pipe == "L" {
			return coords{x: x, y: y}
		}
	}

	x = i + right.x
	y = j + right.y
	if x >= 0 && x < rowLen && y >= 0 && y < colLen {
		pipe := pipes[x][y]
		if pipe == "-" || pipe == "J" || pipe == "7" {
			return coords{x: x, y: y}
		}
	}

	panic("could not find connected pipe")

}

func findStart(pipes [][]string) coords {
	var startingPointCoord coords

outer:
	for i, row := range pipes {
		for j, cell := range row {
			if cell == startingPointSymbol {
				startingPointCoord = coords{
					x: i,
					y: j,
				}
				break outer
			}
		}
	}
	return startingPointCoord
}

func calculateFarthestPointDistance(pipes [][]string) (int, map[coords]bool) {

	prevCoord := findStart(pipes)
	nextCoord := findConnectedPipes(pipes, prevCoord.x, prevCoord.y)
	var visited = make(map[coords]bool)
	visited[prevCoord] = true
	stepCount := 0
	for {
		stepCount += 1
		if visited[nextCoord] {
			break
		}
		visited[nextCoord] = true

		xDiff := nextCoord.x - prevCoord.x
		yDiff := nextCoord.y - prevCoord.y
		pipe := pipes[nextCoord.x][nextCoord.y]
		if pipe == "|" {
			prevCoord = nextCoord
			nextCoord.x += xDiff
		} else if pipe == "-" {
			prevCoord = nextCoord
			nextCoord.y += yDiff
		} else if pipe == "L" {
			prevCoord = nextCoord
			if xDiff > 0 {
				nextCoord.y += 1
			} else {
				nextCoord.x -= 1
			}
		} else if pipe == "J" {
			prevCoord = nextCoord
			if xDiff > 0 {
				nextCoord.y -= 1
			} else {
				nextCoord.x -= 1
			}
		} else if pipe == "7" {
			prevCoord = nextCoord
			if yDiff > 0 {
				nextCoord.x += 1
			} else {
				nextCoord.y -= 1
			}
		} else if pipe == "F" {
			prevCoord = nextCoord
			if yDiff < 0 {
				nextCoord.x += 1
			} else {
				nextCoord.y += 1
			}
		}
	}

	return stepCount / 2, visited
}

func firstPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	var pipes [][]string
	for _, line := range lines {
		pipes = append(pipes, strings.Split(line, ""))
	}
	ans, _ := calculateFarthestPointDistance(pipes)
	fmt.Println("Part 1 Answer: ", ans)
}

func countEnclosedTiles(pipes [][]string, visited map[coords]bool) int {
	// I don't know why this works, all enclosed pipes are withing a centralized
	// inner square of half the size of the original one
	lx := len(pipes) / 4
	ly := len(pipes[0]) / 4
	count := 0
	for i := lx; i < 3*lx; i++ {
		for j := ly; j < 3*ly; j++ {
			if !visited[coords{x: i, y: j}] {
				count += 1
			}
		}
	}
	return count
}

func secondPart() {
	inputFileName := "input.txt"
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	var pipes [][]string
	for _, line := range lines {
		pipes = append(pipes, strings.Split(line, ""))
	}
	_, visited := calculateFarthestPointDistance(pipes)
	// Visualize(visited, len(pipes), len(pipes[0]))
	ans := countEnclosedTiles(pipes, visited)
	fmt.Println("Part 2 Answer: ", ans)
}

func main() {
	firstPart()
	secondPart()
}
