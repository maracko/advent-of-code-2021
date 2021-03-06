//Starting with this problem we use one file for both parts
//Also adding helper funcs, will fully refactor later
package main

import (
	"fmt"

	"github.com/maracko/advent-of-code-2021/helpers/file"
)

type coordinate struct {
	x1, x2, y1, y2 int
}

type Matrix2D [][]int

func main() {
	data, _ := file.ReadFileToSliceOfStrings("data.txt")

	var coords []coordinate
	for _, val := range data {
		c := getCoordinates(val)
		coords = append(coords, c)
	}

	cMat := create2DMatrix(1000)
	cMat2 := create2DMatrix(1000)

	addLinesToMatrix(cMat, coords, false)
	addLinesToMatrix(cMat2, coords, true)

	fmt.Println("Part 1 solution:", getCountOfOverlappingPoints(cMat))
	fmt.Println("Part 2 solution:", getCountOfOverlappingPoints(cMat2))
}

func getCoordinates(line string) coordinate {
	var c coordinate
	fmt.Sscanf(line, "%d,%d -> %d,%d\n", &c.x1, &c.y1, &c.x2, &c.y2)
	return c

}

func create2DMatrix(length int) Matrix2D {
	m := make(Matrix2D, length)
	for i := 0; i < length; i++ {
		m[i] = make([]int, length)
	}
	return m
}

func sortCoordinatesByValue(c *coordinate) {

	if c.x2 < c.x1 {
		c.x1, c.x2 = c.x2, c.x1
	}
	if c.y2 < c.y1 {
		c.y1, c.y2 = c.y2, c.y1
	}
}

func getSortedCoordValue(coord coordinate, axis string) (low int, high int) {
	switch axis {
	case "x":
		if coord.x1 < coord.x2 {
			low = coord.x1
			high = coord.x2
		} else {
			low = coord.x2
			high = coord.x1
		}
	case "y":
		if coord.y1 < coord.y2 {
			low = coord.y1
			high = coord.y2
		} else {
			low = coord.y2
			high = coord.y1
		}
	}
	return
}

func addLinesToMatrix(mat Matrix2D, coords []coordinate, addDiagonal bool) {
	for _, c := range coords {
		if (c.x1 == c.x2) || (c.y1 == c.y2) {
			sortCoordinatesByValue(&c)
			sameCol := c.x1 == c.x2
			sameRow := c.y1 == c.y2
			switch {
			case sameCol:
				column := c.x1
				for row := c.y1; row <= c.y2; row++ {
					mat[row][column]++
				}
			case sameRow:
				row := c.y1
				for column := c.x1; column <= c.x2; column++ {
					mat[row][column]++
				}
			}
		} else if addDiagonal {
			lowX, highX := getSortedCoordValue(c, "x")
			lowY, highY := getSortedCoordValue(c, "y")

			deltaX := highX - lowX
			deltaY := highY - lowY

			if deltaX == deltaY {
				for i := 0; i <= deltaX; i++ {
					j := i
					if c.y1 > c.y2 {
						j = -i
					}
					k := i
					if c.x1 > c.x2 {
						k = -i
					}
					mat[c.y1+j][c.x1+k]++
				}
			}
		}
	}
}

func getCountOfOverlappingPoints(mat Matrix2D) int {
	count := 0
	for _, row := range mat {
		for _, point := range row {
			if point > 1 {
				count++
			}
		}
	}
	return count
}
