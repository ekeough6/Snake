package main

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const WIDTH int32 = 800
const HEIGHT int32 = 600

func main() {
	var direction = 2
	rand.Seed(time.Time.Unix(time.Now()))
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		WIDTH, HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)
	snake := []sdl.Rect{sdl.Rect{10, 10, 10, 10}}
	food := sdl.Rect{rand.Int31n(WIDTH/10) * 10, rand.Int31n(HEIGHT/10) * 10, 10, 10}

	println(food.X, food.Y)

	surface.FillRect(&snake[0], 0xff0000)
	surface.FillRect(&food, 0x00ff00)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				println(t)
				running = false
				break
			}
		}
		states := sdl.GetKeyboardState()
		if states[sdl.SCANCODE_W] == 1 {
			direction = 0
		}
		if states[sdl.SCANCODE_S] == 1 {
			direction = 2
		}
		if states[sdl.SCANCODE_A] == 1 {
			direction = 1
		}
		if states[sdl.SCANCODE_D] == 1 {
			direction = 3
		}

		for i := len(snake) - 1; i > 0; i-- {
			setRect(&snake[i], snake[i-1].X, snake[i-1].Y)
		}

		switch direction {
		case 0:
			moveRect(&snake[0], 0, -snake[0].H)
			break
		case 1:
			moveRect(&snake[0], -snake[0].W, 0)
			break
		case 2:
			moveRect(&snake[0], 0, snake[0].H)
			break
		case 3:
			moveRect(&snake[0], snake[0].W, 0)
			break
		}

		if snake[0].X == food.X && snake[0].Y == food.Y {
			food.X = rand.Int31n(WIDTH/10) * 10
			food.Y = rand.Int31n(HEIGHT/10) * 10
			snake = append(snake, sdl.Rect{snake[len(snake)-1].X, snake[len(snake)-1].Y, 10, 10})
		}

		surface.FillRect(nil, 0)
		for i := range snake {
			surface.FillRect(&snake[i], 0xff0000)
		}
		surface.FillRect(&food, 0x00ff00)
		window.UpdateSurface()
		sdl.Delay(1000 / 15)

	}

}

func moveRect(rect *sdl.Rect, x int32, y int32) {
	rect.Y = rect.Y + y
	rect.X = rect.X + x
	if rect.Y < 0 {
		rect.Y = 0
	}
	if rect.Y > HEIGHT-rect.H {
		rect.Y = HEIGHT - rect.H
	}
	if rect.X < 0 {
		rect.X = 0
	}
	if rect.X > WIDTH-rect.W {
		rect.X = WIDTH - rect.W
	}
}

func setRect(rect *sdl.Rect, x int32, y int32) {
	rect.X = x
	rect.Y = y
}
