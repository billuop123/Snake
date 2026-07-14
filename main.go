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

const (
	dirRight string = "D"
	dirLeft  string = "A"
	dirUp    string = "W"
	dirDown  string = "S"
)

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
	var gridNum int32 = 30
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
	dir := dirRight
	var speed float32 = 0.5
	var gridTimer float32 = 0
	var moveInterval float32 = 0.15
	xRand := rand.Intn(int(gridNum-2)) + 1
	yRand := rand.Intn(int(gridNum-2)) + 1
	gameOver := false
	score := 0
	const increment float32 = 0.25
	headRight := rl.LoadTexture("./headLeft.png")
	tail := rl.LoadTexture("./tail.png")
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
		rl.DrawText(strconv.Itoa(score), 400, 10, 25, rl.Black)
		handleKeyEvents(&dir, &score, &head, &second, &s, &speed, &gridTimer, &gameOver)
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
			if i == 0 {
				handleLoop(headRight, dir, gridW, s, padX, padY, i, "head")
				continue
			} else if i == len(s)-1 {
				handleLoop(tail, dir, gridW, s, padX, padY, i, "tail")
				continue
			}
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
				speed += increment
			}
			score++
		}
		rl.EndDrawing()
	}
}

func drawBox(padX, padY, boxW, boxH int32) {
	green := rl.Color{R: 144, G: 238, B: 144, A: 255}
	rl.DrawRectangle(padX, padY, boxW, boxH, green)
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
	var rectY int32 = 10
	var rectW int32 = 175
	var rectH int32 = 30
	var textFont int32 = 25
	var posX int32 = 630
	var posY int32 = 14
	rl.DrawText("Game is over", textX, textY, textFont, rl.Black)
	rl.DrawRectangle(rectX, rectY, rectW, rectH, rl.Black)
	rl.DrawText("press n", posX, posY, textFont, rl.LightGray)
}

func handleKeyEvents(dir *string, score *int, head *Snake, second *Snake, s *[]Snake, speed *float32, gridTimer *float32, gameOver *bool) {
	switch rl.GetKeyPressed() {
	case rl.KeyS:
		if *dir != dirUp {
			*dir = dirDown
		}
	case rl.KeyW:
		if *dir != dirDown {
			*dir = dirUp
		}
	case rl.KeyD:
		if *dir != dirLeft {
			*dir = dirRight
		}
	case rl.KeyA:
		if *dir != dirRight {
			*dir = dirLeft
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
		*dir = dirRight
		(*s)[0] = *head
		(*s)[1] = *second
		*speed = 0.5
		*gridTimer = 0
		*gameOver = false
		*score = 0
	}
}

func handleLoop(headOrTail rl.Texture2D, dir string, gridW int32, s []Snake, padX, padY int32, i int, somename string) {
	tailDir := headOrTail
	var rotation float32 = 0
	var offsetX float32 = 0
	var offsetY float32 = 0
	if somename == "tail" {
		dir = getTailDirection(s)
	}
	switch dir {
	case dirLeft:
		rotation = 0
		offsetX += float32(gridW)
		offsetY += float32(gridW)
	case dirRight:
		rotation = 180
	case dirUp:
		offsetY += float32(gridW)
		rotation = 90
	case dirDown:
		offsetX += float32(gridW)
		offsetY -= 2
		rotation = -90
	}
	srcRec := rl.Rectangle{
		X:      0,
		Y:      0,
		Height: float32(tailDir.Height),
		Width:  float32(tailDir.Width),
	}
	destRec := rl.Rectangle{
		X:      float32(padX) + s[i].X*float32(gridW) + float32(offsetX),
		Y:      float32(padY) + s[i].Y*float32(gridW) + float32(offsetY),
		Width:  float32(gridW),
		Height: float32(gridW),
	}
	origin := rl.Vector2{X: destRec.Width, Y: destRec.Height}
	rl.DrawTexturePro(tailDir, srcRec, destRec, origin, rotation, rl.White)
}

func getTailDirection(s []Snake) string {
	n := len(s)
	tail := s[n-1]
	prev := s[n-2]
	switch {
	case prev.X < tail.X:
		return dirLeft
	case prev.X > tail.X:
		return dirRight
	case prev.Y > tail.Y:
		return dirDown
	case prev.Y < tail.Y:
		return dirUp
	}
	return dirRight
}
