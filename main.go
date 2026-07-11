package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type coord int

const (
	X coord = iota
	Y
)

type Snake struct {
	X float32
	Y float32
}

func main() {
	var width int32 = 800
	var height int32 = 800
	rl.InitWindow(width, height, "u are a snake")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	var padX int32 = 20
	var padY int32 = 30
	boxW := width - 2*padX
	boxH := boxW
	var gridNum int32 = 10
	gridW := boxH / gridNum
	_ = gridW
	head := Snake{
		X: 2,
		Y: 0,
	}
	second := Snake{
		X: 1,
		Y: 0,
	}
	s := make([]Snake, 2)
	s[0] = head
	s[1] = second
	dir := "D"
	var speed float32 = 0.5
	var gridTimer float32 = 0
	var moveInterval float32 = 0.15
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		drawBox(padX, padY, boxW, boxH)
		switch rl.GetKeyPressed() {
		case rl.KeyS:
			dir = "S"
		case rl.KeyW:
			dir = "W"
		case rl.KeyD:
			dir = "D"
		case rl.KeyA:
			dir = "A"
		}
		gridTimer += rl.GetFrameTime()
		if gridTimer > moveInterval {
			newHead := s[0]
			switch dir {
			case "S":
				newHead.Y += speed
			case "W":
				newHead.Y -= speed
			case "D":
				newHead.X += speed
			case "A":
				newHead.X -= speed
			}
			s[0] = newHead
			s[1] = head
			head = newHead
			gridTimer = 0
		}
		for i := range s {
			if s[i].X != 0 || s[i].Y != 0 {
				rl.DrawRectangle(int32(float32(padX)+s[i].X*float32(gridW)),
					int32(float32(padY)+s[i].Y*float32(gridW)),
					gridW,
					gridW,
					rl.Red)
			}
		}
		rl.EndDrawing()
	}
}

func drawBox(padX, padY, boxW, boxH int32) {
	rl.DrawRectangle(padX, padY, boxW, boxH, rl.LightGray)
}
