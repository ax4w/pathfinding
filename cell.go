package main

import "math"

type (
	cell struct {
		parent                 *cell
		blocked, hit, shortest bool
		x, y, h, g, f, index   int
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

func (c *cell) isInList(list cellList) bool {
	for _, v := range list {
		if v.x == c.x && v.y == c.y {
			return true
		}
	}
	return false
}
