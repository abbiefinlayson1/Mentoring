package main

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
	"strings"
)

type GameObject struct {
	Description string
	MoveX       int
	MoveY       int
	Win         bool
}

func main() {
	// Ask the user for their name and greet them
	fmt.Println("Welcome to the game! Please enter your name:")

	// Read the name input from the user
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello %s! Let's start the adventure.", strings.TrimSpace(name))

	// Assign player's starting position
	x, y := 0, 0

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	// Map setup for rows and columns
	gameMap := [][]string{
		{" ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " "},
	}
	gameMap[2][2] = "G" // goblin
	gameMap[1][4] = "T" // tree
	gameMap[1][1] = "H" // house
	gameMap[2][4] = "C" // castle - winning point!

	// assigning directions to a slice
	directions := []string{"North", "East", "South", "West"}
	dirIndex := 0

	// creating a map of objects for descriptions, move spaces and win for the castle.
	objects := map[string]GameObject{
		"T": {
			Description: "You have reached a tree 🌳. There are apples on it 🍎. You grab one to eat... oh no it's poisonous! Go back 2 spaces.",
			MoveX:       0,
			MoveY:       -2,
		},
		"G": {
			Description: "A Goblin 👹 jumps out and pushes you forward! Move forward one space.",
			MoveX:       1,
			MoveY:       0,
		},
		"H": {
			Description: "You spot a house 🏠 and decide to go in. There is a comfy bed for you to take a nap. You wake up feeling refreshed and move forward 3 spaces",
			MoveX:       3,
			MoveY:       0,
		},
		"C": {
			Description: "You have reached the Castle! Congratulations, you win! 🏆",
			Win:         true,
		},
	}

	// game loop starts here
	for {
		// if player presses Q game ends
		fmt.Println("\nPress a key to continue or 'Q' to quit:")

		char, key, _ := keyboard.GetKey()
		if char == 'q' || char == 'Q' {
			fmt.Println("Goodbye! 👋🏻")
			break
		}

		// Update facing direction first
		if key == keyboard.KeyArrowLeft {
			dirIndex = (dirIndex - 1 + 4) % 4
		}
		if key == keyboard.KeyArrowRight {
			dirIndex = (dirIndex + 1 + 4) % 4
		}
		direction := directions[dirIndex]

		// Movement only for up and down arrows. Facing direction only for left and right to rotate
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

		fmt.Println("Now facing:", direction)

		// Print map with player position
		for rowIndex, row := range gameMap {
			for colIndex, cell := range row {
				if rowIndex == y && colIndex == x {
					fmt.Print("🧝‍♂️")
				} else {
					fmt.Print(cell + " ")
				}
			}
			fmt.Println()
		}

		// Interaction
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
				symbol := gameMap[lookY][lookX]
				obj, exists := objects[symbol]
				if exists {
					fmt.Println(obj.Description)

					newX := x + obj.MoveX
					newY := y + obj.MoveY

					if newX >= 0 && newX < len(gameMap[0]) {
						x = newX
					}
					if newY >= 0 && newY < len(gameMap) {
						y = newY
					}

					if obj.Win {
						fmt.Println("🎉 You reached the castle and won the game!")
						return
					}
				} else {
					fmt.Println("There's nothing in front of you.")
				}
			} else {
				fmt.Println("You have reached the edge of the map!")
			}
		}
		fmt.Printf("You moved to position (%d, %d)\n", x, y)
	}
}
