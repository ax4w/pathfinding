package main

import (
	"sort"
	"time"
)

func solveDijkstra(grid [DIM][DIM]*cell, start, finish *cell) {
	defer func() {
		lock = false
	}()
	resetPathfindingValues(grid)
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
			fromGrid := grid[x][y]
			if fromGrid.isDestination(finish) {
				fromGrid.parent = current
				printTrace(fromGrid)
				return
			}
			if fromGrid.isInList(closeList) {
				continue
			}
			tg := current.g + 1
			if tg < fromGrid.g {
				fromGrid.g = tg
				fromGrid.parent = current
				fromGrid.hit = true
			}
			if !fromGrid.isInList(openList) {
				openList = append(openList, fromGrid)
			}
		}
		time.Sleep(20 * time.Millisecond)
	}
	println("Not found")
}
