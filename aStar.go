package main

import (
	"math"
	"sort"
	"time"
)

type (
	cell struct {
		parent                 *cell
		blocked, hit, shortest bool
		x, y, h, g, f          int
	}
)

func (c *cell) getNeighbours(grid [DIM][DIM]*cell) []*cell {
	var res []*cell
	var u, r, d, l bool
	if isIndexValid(c.x, c.y-1) && !grid[c.x][c.y-1].blocked {
		res = append(res, grid[c.x][c.y-1])
		u = true
	}
	if isIndexValid(c.x+1, c.y) && !grid[c.x+1][c.y].blocked {
		res = append(res, grid[c.x+1][c.y])
		r = true
	}
	if isIndexValid(c.x, c.y+1) && !grid[c.x][c.y+1].blocked {
		res = append(res, grid[c.x][c.y+1])
		d = true
	}
	if isIndexValid(c.x-1, c.y) && !grid[c.x-1][c.y].blocked {
		res = append(res, grid[c.x-1][c.y])
		l = true
	}
	//top left
	if l && u && isIndexValid(c.x-1, c.y-1) && !grid[c.x-1][c.y-1].blocked {
		res = append(res, grid[c.x-1][c.y-1])
	}
	// top right
	if r && u && isIndexValid(c.x+1, c.y-1) && !grid[c.x+1][c.y-1].blocked {
		res = append(res, grid[c.x+1][c.y-1])
	}
	// bottom right
	if d && r && isIndexValid(c.x+1, c.y+1) && !grid[c.x+1][c.y+1].blocked {
		res = append(res, grid[c.x+1][c.y+1])
	}
	// bottom left
	if d && l && isIndexValid(c.x-1, c.y+1) && !grid[c.x-1][c.y+1].blocked {
		res = append(res, grid[c.x-1][c.y+1])
	}
	return res
}

func (c *cell) isDestination(finish *cell) bool {
	return c.x == finish.x && c.y == finish.y
}

// euclidean formular
func (c *cell) getHScore(finish *cell) int {
	x := math.Pow(float64(c.x-finish.x), 2)
	y := math.Pow(float64(c.y-finish.y), 2)
	return int(math.Sqrt(x + y))
}

func isIndexValid(i, j int) bool {
	return i >= 0 && j >= 0 && i < DIM && j < DIM
}

func (c *cell) isInList(list cellList) bool {
	for _, v := range list {
		if v.x == c.x && v.y == c.y {
			return true
		}
	}
	return false
}

func printTrace(c *cell) {
	for c != nil {
		c.shortest = true
		c = c.parent
	}
}

func solve(grid [DIM][DIM]*cell, start, finish *cell) {
	if start == nil || finish == nil {
		return
	}
	if start == finish {
		return
	}
	defer func() {
		lock = false
	}()
	for i := 0; i < DIM; i++ {
		for j := 0; j < DIM; j++ {
			grid[i][j].g = INF
			grid[i][j].f = INF
			grid[i][j].h = INF
			grid[i][j].shortest = false
			grid[i][j].hit = false
			grid[i][j].parent = nil
		}
	}
	var openList, closeList cellList
	start.h = 0
	start.f = 0
	start.g = 0
	openList = append(openList, start)
	for len(openList) > 0 {
		sort.Sort(openList)
		current := openList[0]
		openList = openList[1:]
		closeList = append(closeList, current)
		for _, v := range current.getNeighbours(grid) {
			x := v.x
			y := v.y
			if !isIndexValid(x, y) {
				continue
			}
			fromGrid := grid[x][y]
			if fromGrid.isDestination(finish) {
				fromGrid.parent = current
				printTrace(fromGrid)
				return
			}
			if fromGrid.isInList(closeList) {
				continue
			}
			gNew := current.g + 1
			hNew := fromGrid.getHScore(finish)
			fNew := gNew + hNew
			if fromGrid.f == INF || fromGrid.f > fNew {
				fromGrid.g = gNew
				fromGrid.h = hNew
				fromGrid.f = fNew
				fromGrid.parent = current
				fromGrid.hit = true
				openList = append(openList, fromGrid)
			}
		}
		time.Sleep(20 * time.Millisecond)
	}
	println("Not found")
}
