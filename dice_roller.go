package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	sim "example.com/gottrpg/v2/simulation"
)

// rollDiceInteractive accepts user input to roll dice
func rollDiceInteractive() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("=== Dice Roller ===")
	fmt.Println("Enter dice commands (e.g., '3d6', 'd20', '2d8') or 'quit' to exit:")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(strings.ToLower(scanner.Text()))
		if input == "quit" || input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		result := parseDiceCommand(input)
		if result != "" {
			fmt.Println(result)
		}
	}
}

// parseDiceCommand parses dice notation and returns the roll result
func parseDiceCommand(input string) string {
	// Handle single die notation (e.g., "d6", "d20")
	if strings.HasPrefix(input, "d") {
		sidesStr := input[1:]
		sides, err := strconv.Atoi(sidesStr)
		if err != nil {
			return "Invalid dice notation. Use format like 'd6' or '3d8'"
		}

		// check for valid die
		if err := validatePolyhedral(sides); err != nil {
			return err.Error()
		}

		return fmt.Sprintf("%d", sim.DiceRoll(sides, 1))
	}

	// Handle multiple dice notation (e.g., "3d6", "2d8")
	parts := strings.Split(input, "d")
	if len(parts) != 2 {
		return "Invalid dice notation. Use format like 'd6' or '3d8'"
	}

	quantity, err1 := strconv.Atoi(parts[0])
	sides, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil {
		return "Invalid dice notation. Use format like 'd6' or '3d8'"
	}

	if quantity <= 0 || sides <= 0 {
		return "Quantity and sides must be positive numbers"
	}

	// check for valid die
	if err := validatePolyhedral(sides); err != nil {
		return err.Error()
	}

	return fmt.Sprintf("%d", sim.DiceRoll(sides, quantity))
}

func validatePolyhedral(sides int) error {
	switch sides {
	case 4, 6, 8, 10, 12, 20:
		return nil
	default:
		break
	}
	return errors.New("select a polyhedral with 4, 6, 8, 10, 12, or 20 sides")
}