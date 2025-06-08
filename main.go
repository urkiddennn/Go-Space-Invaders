package main

import (
	"fmt"
	"log"
	"time"

	"github.com/eiannone/keyboard"
)

const (
	GREEN       = "\033[32m"
	RED         = "\033[31m"
	GRID_WIDTH  = 20
	GRID_HEIGHT = 20
	SQUARE_CHAR = GREEN + "■"
	APPLE_CHAR  = RED + "■"
	EMPTY_CHAR  = " "
	FRAME_RATE  = 250 * time.Millisecond
)

type (
	Grid   [][]string
	Bullet struct {
		x int
		y int
	}
)

type Player struct {
	x          int
	y          int
	shotBullet bool
	isCollided bool
}

type Direction struct {
	dx, dy int
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func newGrid() Grid {
	grid := make(Grid, GRID_HEIGHT)
	for i := range grid {
		grid[i] = make([]string, GRID_WIDTH)
		for j := range grid[i] {
			grid[i][j] = EMPTY_CHAR
		}
	}
	return grid
}

func drawGrid(grid Grid, player Player, bullet Bullet) {
	clearScreen()

	// Create a fresh grid for drawing
	tempGrid := newGrid()

	if player.x < GRID_WIDTH && player.x >= 0 {
		tempGrid[player.y][player.x] = SQUARE_CHAR
	}
	for _, row := range tempGrid {
		for _, char := range row {
			fmt.Printf("%s", char)
		}
		fmt.Println()
	}
}

// Shot Bullter function
func shotBullet(player Player) Bullet {
	for {
		bullet := Bullet{x: player.x, y: player.y + 1}
		// collsion := false
		// visible := false

		return bullet

	}
}

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	grid := newGrid()

	player := Player{x: GRID_WIDTH / 2, y: GRID_HEIGHT / 2, shotBullet: false, isCollided: false}
	bullet := shotBullet(player)
	dir := Direction{0, 0}

	keyEvents, err := keyboard.GetKeys(10)
	if err != nil {
		fmt.Println("Error getting keys:", err)
		return
	}

	for {
		select {
		case event := <-keyEvents:
			if event.Err != nil {
				fmt.Println("Keyboard event error:", event.Err)
				return

			}
			switch event.Rune {
			case 'a', 'A':
				if dir.dx != 1 {
					dir = Direction{-1, 0}
				}
			case 'd', 'D':
				if dir.dx != -1 {
					dir = Direction{1, 0}
				}
			case rune(keyboard.KeySpace):
				player.shotBullet = true
				fmt.Println("Space Key pressed!!")

			case 'q':
				return
			}

		default:
			player.x = dir.dx
			drawGrid(grid, player, bullet)
			time.Sleep(FRAME_RATE)
		}
	}
}
