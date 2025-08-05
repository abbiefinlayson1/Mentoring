package main

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
	"strings"
)

func main() {
	// Ask the user for their name and greet them
	fmt.Println("Welcome to the game! Please enter your name:")

	// Read the input from the user
	reader := bufio.NewReader(os.Stdin)
	// Stores the input, ignores any error and stops when the user pressed Enter
	name, _ := reader.ReadString('\n')
	// Print a greeting message
	fmt.Printf("Hello %s! Let's start the adventure.", strings.TrimSpace(name))

	// Assign player's starting position
	x, y := 0, 0

	// keyboard input handling
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	// creating the game map
	gameMap := [][]string{
		{" ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " "},
	}
	gameMap[2][2] = "G" // Goblin position here - player can move to this position to interact with the Goblin
	gameMap[1][4] = "T" // Tree position here
	gameMap[1][1] = "H" // House position here
	gameMap[2][4] = "C" // Castle position here
	// Print the game map
	fmt.Println("\nGame Map:")
	// creating a slice of directions
	directions := []string{"North", "East", "South", "West"}
	// the direction index starts at 0 = North
	dirIndex := 0

	for {
		fmt.Println("\nPress a key to continue or 'Q' to quit:")

		char, key, _ := keyboard.GetKey()
		if char == 'q' || char == 'Q' {
			fmt.Println("Goodbye! 👋🏻")
			break
		}

		direction := directions[dirIndex]

		// change facing position
		// change position based on arrow keys
		if key == keyboard.KeyArrowUp {
			if direction == "North" && y > 0 {
				y--
			}
			if direction == "South" && y < len(gameMap)-1 {
				y++
			}
			if direction == "East" && x < len(gameMap[0])-1 {
				x++
			}
			if direction == "West" && x > 0 {
				x--
			}

		}
		if key == keyboard.KeyArrowDown {
			if direction == "North" && y < len(gameMap)-1 {
				y++
			}
			if direction == "South" && y > 0 {
				y--
			}
			if direction == "East" && x > 0 {
				x--
			}
			if direction == "West" && x < len(gameMap[0])-1 {
				x++
			}
		}
		if key == keyboard.KeyArrowLeft {
			dirIndex = (dirIndex - 1 + 4) % 4
		}
		if key == keyboard.KeyArrowRight {
			dirIndex = (dirIndex + 1 + 4) % 4
		}
		fmt.Println("Now facing:", direction)

		// Print the game map with the player's position
		for rowIndex, row := range gameMap {
			for colIndex, cell := range row {
				if rowIndex == y && colIndex == x {
					fmt.Print("P ")
				} else {
					fmt.Print(cell + " ")
				}
			}
			fmt.Println()
		}
		// pressing D will describe what is in front of the player
		if char == 'D' || char == 'd' {
			lookX := x
			lookY := y
			if direction == "North" {
				lookY--
			}
			if direction == "South" {
				lookY++
			}
			if direction == "East" {
				lookX++
			}
			if direction == "West" {
				lookX--
			}

			if lookX >= 0 && lookX < len(gameMap[0]) && lookY >= 0 && lookY < len(gameMap) {
				object := gameMap[lookY][lookX]
				if object == " " {
					fmt.Println("There is nothing in front of you!")
				} else {
					fmt.Printf("You see a %s in front of you \n", object)
				}
			} else {
				fmt.Println("You have reached the end of the world, please go back!")
			}
		}
		fmt.Printf("You moved to position (%d, %d)\n", x, y)
	}
}
