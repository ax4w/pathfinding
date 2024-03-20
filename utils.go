package main

func printTrace(c *cell) {
	for c != nil {
		c.shortest = true
		c = c.parent
	}
}

func isIndexValid(i, j int) bool {
	return i >= 0 && j >= 0 && i < DIM && j < DIM
}

func resetPathfindingValues(grid [DIM][DIM]*cell) {
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
}
