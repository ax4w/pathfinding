package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	DIM      = 50
	WIN      = 1250
	INF      = math.MaxInt
	CONTROLS = 250
	BORDER   = 3
)

var (
	start  *cell
	finish *cell
	RECT   = float32(20)
	lock   = false
)

func mouseOverlapsField(mx, my, rx, ry float32) bool {
	return mx >= rx && my >= ry && mx <= rx+RECT && my <= ry+RECT
}

func drawControls() {
	rl.DrawRectangle(10, WIN+(CONTROLS/6)-20, 30, 30, rl.Red)
	rl.DrawText("press s to set the starting point on the current mouse position", 50, WIN+(CONTROLS/6)-20, 30, rl.White)
	rl.DrawRectangle(10, WIN+(CONTROLS/3), 30, 30, rl.Green)
	rl.DrawText("press t to set the target point on the current mouse position", 50, WIN+(CONTROLS/3), 30, rl.White)
	rl.DrawRectangle(10, WIN+(CONTROLS/2)+20, 30, 30, rl.Gray)
	rl.DrawText("hold the l/r mouse button and drag over the field to draw / delete barriers", 50, WIN+(CONTROLS/2)+20, 30, rl.White)
	rl.DrawText("press c to clear the field", 50, WIN+CONTROLS-50, 30, rl.White)
}

func main() {
	rl.InitWindow(WIN, WIN+CONTROLS, "A*")
	RECT = WIN / DIM
	defer rl.CloseWindow()
	grid := [DIM][DIM]*cell{}
	for i := 0; i < DIM; i++ {
		for j := 0; j < DIM; j++ {
			grid[i][j] = &cell{x: i, y: j, g: INF, f: INF}
		}
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		for i := 0; i < DIM; i++ {
			for j := 0; j < DIM; j++ {
				c := grid[i][j]
				color := rl.White
				if c == start {
					color = rl.Red
				} else if c == finish {
					color = rl.Green
				} else if c.blocked {
					color = rl.Gray
				} else if c.hit {
					color = rl.Black
				}
				if c.shortest {
					color = rl.Yellow
				}
				rl.DrawRectangle(int32(RECT*float32(c.x))+BORDER, int32(RECT*float32(c.y))+BORDER, int32(RECT-BORDER), int32(RECT-BORDER), color)
			}
		}
		mouseVec := rl.GetMousePosition()
		if rl.IsKeyPressed(rl.KeyS) && !lock {
			for i := 0; i < DIM; i++ {
				for j := 0; j < DIM; j++ {
					c := grid[i][j]
					if mouseOverlapsField(mouseVec.X, mouseVec.Y, float32(RECT*float32(c.x)), float32(RECT*float32(c.y))) {
						start = c
					}
				}
			}
		} else if rl.IsKeyPressed(rl.KeyT) && !lock {
			for i := 0; i < DIM; i++ {
				for j := 0; j < DIM; j++ {
					c := grid[i][j]
					if mouseOverlapsField(mouseVec.X, mouseVec.Y, float32(RECT*float32(c.x)), float32(RECT*float32(c.y))) {
						finish = c
					}
				}
			}

		} else if rl.IsMouseButtonDown(rl.MouseLeftButton) || rl.IsMouseButtonDown(rl.MouseRightButton) && !lock {
			for i := 0; i < DIM; i++ {
				for j := 0; j < DIM; j++ {
					c := grid[i][j]
					if mouseOverlapsField(mouseVec.X, mouseVec.Y, float32(RECT*float32(c.x)), float32(RECT*float32(c.y))) {
						c.blocked = rl.IsMouseButtonDown(rl.MouseLeftButton)

					}
				}
			}
		} else if rl.IsKeyPressed(rl.KeySpace) && !lock {
			lock = true
			go solve(grid, start, finish)
		} else if rl.IsKeyPressed(rl.KeyC) && !lock {
			//clear
			start, finish = nil, nil
			for i := 0; i < DIM; i++ {
				for j := 0; j < DIM; j++ {
					grid[i][j].g = INF
					grid[i][j].f = INF
					grid[i][j].h = INF
					grid[i][j].blocked = false
					grid[i][j].shortest = false
					grid[i][j].hit = false
					grid[i][j].parent = nil
					//grid[i][j] = &cell{x: i, y: j, g: INF, f: INF}
				}
			}
		}
		drawControls()
		rl.EndDrawing()
	}
}
