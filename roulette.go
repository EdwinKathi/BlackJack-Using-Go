package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var bankroll float64
	fmt.Print("Enter your bankroll amount: $")
	fmt.Scanln(&bankroll)

	for bankroll > 0 {
		fmt.Println("-----")
		fmt.Println("Menu:")
		fmt.Println("1. Bet on red or black")
		fmt.Println("2. Bet on high or low")
		fmt.Println("3. Bet on a specific number")
		fmt.Println("4. Quit the program")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			bankroll = betOnRedOrBlack(bankroll)
		case 2:
			bankroll = betOnHighOrLow(bankroll)
		case 3:
			bankroll = betOnSpecificNumber(bankroll)
		case 4:
			fmt.Println("Quitting the program...")
			return
		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}

	fmt.Println("You ran out of money! Game over.")
}

func betOnRedOrBlack(bankroll float64) float64 {
	var betAmount float64
	fmt.Print("Enter your bet amount: $")
	fmt.Scanln(&betAmount)

	if betAmount > bankroll {
		fmt.Println("Insufficient funds!")
		return bankroll
	}

	var colorChoice int
	fmt.Print("Enter your color choice (1 for red, 2 for black): ")
	fmt.Scanln(&colorChoice)

	spinResult := spinRouletteWheel()

	switch spinResult {
	case "red":
		if colorChoice == 1 {
			bankroll += betAmount
			fmt.Println("Congratulations! You won.")
		} else {
			bankroll -= betAmount
			fmt.Println("Sorry! You lost.")
		}
	case "black":
		if colorChoice == 2 {
			bankroll += betAmount
			fmt.Println("Congratulations! You won.")
		} else {
			bankroll -= betAmount
			fmt.Println("Sorry! You lost.")
		}
	}

	fmt.Printf("Your new bankroll: $%.2f\n", bankroll)
	return bankroll
}

func betOnHighOrLow(bankroll float64) float64 {
	var betAmount float64
	fmt.Print("Enter your bet amount: $")
	fmt.Scanln(&betAmount)

	if betAmount > bankroll {
		fmt.Println("Insufficient funds!")
		return bankroll
	}

	var numberChoice int
	fmt.Print("Enter your number choice (1 for high, 2 for low): ")
	fmt.Scanln(&numberChoice)

	spinResult := spinRouletteWheel()
	number := getRouletteNumber(spinResult)

	switch {
	case numberChoice == 1 && number >= 19 && number <= 36:
		bankroll += betAmount
		fmt.Println("Congratulations! You won.")
	case numberChoice == 2 && number >= 1 && number <= 18:
		bankroll += betAmount
		fmt.Println("Congratulations! You won.")
	default:
		bankroll -= betAmount
		fmt.Println("Sorry! You lost.")
	}

	fmt.Printf("Your new bankroll: $%.2f\n", bankroll)
	return bankroll
}

func betOnSpecificNumber(bankroll float64) float64 {
	var betAmount float64
	fmt.Print("Enter your bet amount: $")
	fmt.Scanln(&betAmount)

	if betAmount > bankroll {
		fmt.Println("Insufficient funds!")
		return bankroll
	}

	var numberChoice int
	fmt.Print("Enter your number choice (0-36): ")
	fmt.Scanln(&numberChoice)

	spinResult := spinRouletteWheel()
	number := getRouletteNumber(spinResult)

	if numberChoice == number {
		winnings := 35 * betAmount
		bankroll += betAmount + winnings
		fmt.Println("Congratulations! You won.")
	} else {
		bankroll -= betAmount
		fmt.Println("Sorry! You lost.")
	}

	fmt.Printf("Your new bankroll: $%.2f\n", bankroll)
	return bankroll
}

func spinRouletteWheel() string {
	colors := []string{"red", "black"}
	return colors[rand.Intn(len(colors))]
}

func getRouletteNumber(color string) int {
	if color == "red" {
		return rand.Intn(18) + 1
	}
	return rand.Intn(18) + 19
}
