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
	Player struct {
		x          int
		y          int
		shotBullet bool
		isCollided bool
	}
	Direction struct {
		dx, dy int
	}
)

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

	// Draw player
	if player.x >= 0 && player.x < GRID_WIDTH && player.y >= 0 && player.y < GRID_HEIGHT {
		tempGrid[player.y][player.x] = SQUARE_CHAR
	}

	// Draw bullet (not fully implemented, but included for consistency)
	if bullet.x >= 0 && bullet.x < GRID_WIDTH && bullet.y >= 0 && bullet.y < GRID_HEIGHT && player.shotBullet {
		tempGrid[bullet.y][bullet.x] = APPLE_CHAR
	}

	// Print the grid
	for _, row := range tempGrid {
		for _, char := range row {
			fmt.Printf("%s", char)
		}
		fmt.Println()
	}
}

// Shot Bullet function (unchanged from original)
func shotBullet(player Player) Bullet {
	for {
		bullet := Bullet{x: player.x, y: player.y + 1}
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
	//	dir := Direction{0, 0} // Kept for compatibility, but not used for movement

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
				// Move left one space if within bounds
				if player.x > 0 {
					player.x -= 1
					fmt.Println("Move left")
				}
			case 'd', 'D':
				// Move right one space if within bounds
				if player.x < GRID_WIDTH-1 {
					player.x += 1
					fmt.Println("Move right")
				}
			case rune(keyboard.KeySpace):
				player.shotBullet = true
				bullet = shotBullet(player) // Update bullet position
				fmt.Println("Space Key pressed!!")
			case 'q':
				return
			}

		default:
			// Redraw grid without updating position
			drawGrid(grid, player, bullet)
			time.Sleep(FRAME_RATE)
		}
	}
}
