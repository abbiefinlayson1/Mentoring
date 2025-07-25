package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	// Ask the user for their name and greet them
	fmt.Println("Welcome to the game! Please enter your name:")

	// Read the input from the user
	reader := bufio.NewReader(os.Stdin)
	// Stores the input, ignores any error and stops when the user pressed Enter
	name, _ := reader.ReadString('\n')
	// Print a greeting message
	fmt.Printf("Hello %s! Let's start the adventure.", strings.TrimSpace(name))
}
