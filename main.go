package main

import (
	"math/rand"
	"strconv"

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
	xRand := rand.Intn(int(gridNum-2)) + 1
	yRand := rand.Intn(int(gridNum-2)) + 1
	gameOver := false
	score := 0
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		drawBox(padX, padY, boxW, boxH)
		rl.DrawRectangle(int32(float32(padX)+float32(xRand)*float32(gridW)),
			int32(float32(padY)+float32(yRand)*float32(gridW)),
			gridW/2,
			gridW/2,
			rl.Yellow)
		if gameOver {
			gameOverContent()
		}
		rl.DrawText(strconv.Itoa(score), 400, 10, 25, rl.Green)
		handleKeyEvents(&dir, &head, &second, &s, &speed, &gridTimer, &gameOver)
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
			if checkSelfCollision(newHead, s) || checkBoundary(newHead, gridNum) {
				speed = 0
				gameOver = true
				if checkBoundary(newHead, gridNum) {
					if newHead.X < 0 {
						newHead.X = 0
					} else if newHead.X > float32(gridNum-1) {
						newHead.X = float32(gridNum - 1)
					} else if newHead.Y < 0 {
						newHead.Y = 0
					} else {
						newHead.Y = float32(gridNum - 1)
					}
				}
				s = append([]Snake{newHead}, s...)
			} else {
				s = append([]Snake{newHead}, s...)
				s = s[:len(s)-1]
				head = newHead
			}
			gridTimer = 0
		}
		for i := range s {
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
			xRand = rand.Intn(int(gridNum-2)) + 1
			yRand = rand.Intn(int(gridNum-2)) + 1
			s = append(s, head)
			if speed < 1 {
				speed += 0.25
			}
			score++
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

func checkBoundary(newHead Snake, gridNum int32) bool {
	if newHead.X < 0 || newHead.Y < 0 || newHead.X > float32(gridNum-1) || newHead.Y > float32(gridNum-1) {
		return true
	}
	return false
}

func gameOverContent() {
	var textX int32 = 50
	var textY int32 = 10
	var rectX int32 = 600
	var rectY int32 = 175
	var rectW int32 = 175
	var rectH int32 = 30
	var textFont int32 = 25
	var posX int32 = 630
	var posY int32 = 14
	rl.DrawText("Game is over", textX, textY, textFont, rl.Black)
	rl.DrawRectangle(rectX, rectY, rectW, rectH, rl.Black)
	rl.DrawText("press n", posX, posY, textFont, rl.LightGray)
}

func handleKeyEvents(dir *string, head *Snake, second *Snake, s *[]Snake, speed *float32, gridTimer *float32, gameOver *bool) {
	switch rl.GetKeyPressed() {
	case rl.KeyS:
		if *dir != "W" {
			*dir = "S"
		}
	case rl.KeyW:
		if *dir != "S" {
			*dir = "W"
		}
	case rl.KeyD:
		if *dir != "A" {
			*dir = "D"
		}
	case rl.KeyA:
		if *dir != "D" {
			*dir = "A"
		}
	case rl.KeyN:
		*s = make([]Snake, 2)
		*head = Snake{
			X: 2,
			Y: 0,
		}
		*second = Snake{
			X: 1,
			Y: 0,
		}
		*dir = "D"
		(*s)[0] = *head
		(*s)[1] = *second
		*speed = 0.5
		*gridTimer = 0
		*gameOver = false
	}
}
