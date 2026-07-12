package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
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
	xRand := rand.Intn(10)
	yRand := rand.Intn(10)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		drawBox(padX, padY, boxW, boxH)
		rl.DrawRectangle(int32(float32(padX)+float32(xRand)*float32(gridW)),
			int32(float32(padY)+float32(yRand)*float32(gridW)),
			gridW/2,
			gridW/2,
			rl.Yellow)
		switch rl.GetKeyPressed() {
		case rl.KeyS:
			if dir != "W" {
				dir = "S"
			}
		case rl.KeyW:
			if dir != "S" {
				dir = "W"
			}
		case rl.KeyD:
			if dir != "A" {
				dir = "D"
			}
		case rl.KeyA:
			if dir != "D" {
				dir = "A"
			}
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
			gridTimer = 0
			if checkSelfCollision(newHead, s) || checkBoundary(newHead) {
				speed = 0
			} else {
				s = append([]Snake{newHead}, s...)
				s = s[:len(s)-1]
				head = newHead
			}
		}
		for i := range s {
			if s[i].X != 0 || s[i].Y != 0 {
				rl.DrawRectangle(int32(float32(padX)+s[i].X*float32(gridW)),
					int32(float32(padY)+s[i].Y*float32(gridW)),
					gridW,
					gridW,
					rl.Red)
			}
			if (s[0].X*float32(gridW) <= float32(xRand)*float32(gridW)) &&
				s[0].X*float32(gridW)+float32(gridW) >= float32(xRand)*float32(gridW)+float32(gridW/2) &&
				s[0].Y*float32(gridW) <= float32(yRand)*float32(gridW) &&
				s[0].Y*float32(gridW)+float32(gridW) >= float32(yRand)*float32(gridW)+float32(gridW/2) {
				xRand = rand.Intn(10)
				yRand = rand.Intn(10)
				s = append(s, head)
				if speed < 1 {
					speed += 0.5
				}
			}
		}
		rl.EndDrawing()
	}
}

func drawBox(padX, padY, boxW, boxH int32) {
	rl.DrawRectangle(padX, padY, boxW, boxH, rl.LightGray)
}

func checkSelfCollision(newHead Snake, s []Snake) bool {
	for i := 1; i < len(s); i++ {
		if newHead.X == s[i].X && newHead.Y == s[i].Y {
			return true
		}
	}
	return false
}

func checkBoundary(newHead Snake) bool {
	if newHead.X < 0 || newHead.Y < 0 || newHead.X >= 9.5 || newHead.Y >= 9.5 {
		return true
	}
	return false
}
