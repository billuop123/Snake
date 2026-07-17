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

type gridMetrics struct {
	padX      int32
	padY      int32
	boxW      int32
	cellCount int32
	cellSize  int32
}

func (g *gridMetrics) init(width int32) {
	g.boxW = width - 2*g.padX
	g.cellSize = g.boxW / g.cellCount
}

func (g gridMetrics) drawBox() {
	green := rl.Color{R: 144, G: 238, B: 144, A: 255}
	rl.DrawRectangle(g.padX, g.padY, g.boxW, g.boxW, green)
}

func (g gridMetrics) checkBoundary(newHead Snake) bool {
	if newHead.X < 0 ||
		newHead.Y < 0 ||
		newHead.X > float32(g.cellCount-1) ||
		newHead.Y > float32(g.cellCount-1) {
		return true
	}
	return false
}

func (g gridMetrics) eatenOrNot(s []Snake, foodXrand, foodYrand int) bool {
	return (s[0].X*float32(g.cellSize) <= float32(foodXrand)*float32(g.cellSize)) &&
		s[0].X*float32(g.cellSize)+float32(g.cellSize) >= float32(foodXrand)*float32(g.cellSize)+
			float32(g.cellSize/2) &&
		s[0].Y*float32(g.cellSize) <= float32(foodYrand)*float32(g.cellSize) &&
		s[0].Y*float32(g.cellSize)+float32(g.cellSize) >= float32(foodYrand)*float32(g.cellSize)+
			float32(g.cellSize/2)
}

type motionMetrics struct {
	speed        float32
	moveTimer    float32
	moveInterval float32
	increment    float32
}
type scoreMetrics struct {
	score    int32
	posX     int32
	posY     int32
	fontSize int32
	color    rl.Color
}

func main() {
	var width int32 = 800
	var height int32 = 800
	startingX := 2
	startingY := 2
	rl.InitWindow(width, height, "u are a snake")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	g := gridMetrics{
		padX:      20,
		padY:      30,
		cellCount: 30,
	}
	g.init(width)
	head := Snake{
		X: float32(startingX),
		Y: float32(startingY),
	}
	s := make([]Snake, 2)
	s[0] = head
	dir := dirRight
	m := motionMetrics{
		speed:        0.5,
		moveTimer:    0,
		moveInterval: 0.15,
		increment:    0.25,
	}
	foodXrand := rand.Intn(int(g.cellCount-2)) + 1
	foodYrand := rand.Intn(int(g.cellCount-2)) + 1
	gameOver := false
	headLeft := rl.LoadTexture("./headLeft.png")
	tail := rl.LoadTexture("./tail.png")
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		g.drawBox()
		rl.DrawRectangle(int32(float32(g.padX)+float32(foodXrand)*float32(g.cellSize)),
			int32(float32(g.padY)+float32(foodYrand)*float32(g.cellSize)),
			g.cellSize/2,
			g.cellSize/2,
			rl.Yellow)
		if gameOver {
			gameOverContent()
		}
		sMet := scoreMetrics{
			score:    0,
			posX:     400,
			posY:     10,
			fontSize: 25,
			color:    rl.Black,
		}
		rl.DrawText(strconv.Itoa(int(sMet.score)), sMet.posX, sMet.posY,
			sMet.fontSize, rl.Black)
		handleKeyEvents(&dir, &sMet, &head, &s, &m, &gameOver)
		m.moveTimer += rl.GetFrameTime()
		if m.moveTimer > m.moveInterval {
			newHead := s[0]
			switch dir {
			case dirDown:
				newHead.Y += m.speed
			case dirUp:
				newHead.Y -= m.speed
			case dirRight:
				newHead.X += m.speed
			case dirLeft:
				newHead.X -= m.speed
			}
			if checkSelfCollision(newHead, s) || g.checkBoundary(newHead) {
				m.speed = 0
				gameOver = true
				if g.checkBoundary(newHead) {
					switch {
					case newHead.X < 0:
						newHead.X = 0
					case newHead.X > float32(g.cellCount-1):
						newHead.X = float32(g.cellCount - 1)
					case newHead.Y < 0:
						newHead.Y = 0
					default:
						newHead.Y = float32(g.cellCount - 1)
					}
				}
				s = append([]Snake{newHead}, s...)
			} else {
				s = append([]Snake{newHead}, s...)
				s = s[:len(s)-1]
				head = newHead
			}
			m.moveTimer = 0
		}
		for i := range s {
			if i == 0 {
				g.handleLoop(headLeft, dir, s, i, "head")
				continue
			} else if i == len(s)-1 {
				g.handleLoop(tail, dir, s, i, "tail")
				continue
			}
			rl.DrawRectangle(int32(float32(g.padX)+s[i].X*float32(g.cellSize)),
				int32(float32(g.padY)+s[i].Y*float32(g.cellSize)),
				g.cellSize,
				g.cellSize,
				rl.Red)
		}
		if g.eatenOrNot(s, foodXrand, foodYrand) {
			foodXrand = rand.Intn(int(g.cellCount-2)) + 1
			foodYrand = rand.Intn(int(g.cellCount-2)) + 1
			s = append(s, head)
			if m.speed < 1 {
				m.speed += m.increment
			}
			sMet.score++
		}
		rl.EndDrawing()
	}
}

func checkSelfCollision(newHead Snake, s []Snake) bool {
	// TODO:
	for i := 1; i < len(s); i++ {
		if newHead.X == s[i].X && newHead.Y == s[i].Y {
			return true
		}
	}
	return false
}

type gameOverMets struct {
	textX    int32
	textY    int32
	rectX    int32
	rectY    int32
	rectW    int32
	rectH    int32
	textFont int32
	posX     int32
	posY     int32
}

func gameOverContent() {
	gOver := gameOverMets{
		textX:    50,
		textY:    10,
		rectX:    600,
		rectY:    10,
		rectW:    175,
		rectH:    30,
		textFont: 25,
		posX:     630,
		posY:     14,
	}
	rl.DrawText("Game is over", gOver.textX, gOver.textY,
		gOver.textFont, rl.Black)
	rl.DrawRectangle(gOver.rectX, gOver.rectY, gOver.rectW, gOver.rectH, rl.Black)
	rl.DrawText("press n", gOver.posX, gOver.posY, gOver.textFont, rl.LightGray)
}

func handleKeyEvents(dir *string, sMet *scoreMetrics, head *Snake, s *[]Snake, m *motionMetrics, gameOver *bool) {
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
		*dir = dirRight
		(*s)[0] = *head
		m.speed = 0.5
		m.moveTimer = 0
		*gameOver = false
		sMet.score = 0
	}
}

func (g gridMetrics) handleLoop(headOrTail rl.Texture2D, dir string, s []Snake, i int, somename string) {
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
		offsetX += float32(g.cellSize)
		offsetY += float32(g.cellSize)
	case dirRight:
		rotation = 180
	case dirUp:
		offsetY += float32(g.cellSize)
		rotation = 90
	case dirDown:
		offsetX += float32(g.cellSize)
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
		X:      float32(g.padX) + s[i].X*float32(g.cellSize) + float32(offsetX),
		Y:      float32(g.padY) + s[i].Y*float32(g.cellSize) + float32(offsetY),
		Width:  float32(g.cellSize),
		Height: float32(g.cellSize),
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
