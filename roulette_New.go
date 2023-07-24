package main

import (
	"fmt"
)

func Bet_amount() int {
	var bet int
	fmt.Print("Enter Bet Amount: ")
	fmt.Scan(&bet)
	if bet <= 0 {
		fmt.Println("Enter a bet amount greater than zero")
		return Bet_amount()
	}
	return bet
}

func main() {
	var bankroll int
	fmt.Print("Enter your initial amount: ")
	fmt.Scan(&bankroll)

	for bankroll > 0 {
		fmt.Println("1. Bet on red or black")
		fmt.Println("2. Bet on high or low")
		fmt.Println("3. Bet on a specific number")
		fmt.Println("4. Quit")

		var choice int
		fmt.Print("Enter your choice (1, 2, 3, 4): ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			bet := Bet_amount()
			fmt.Print("Enter '1' for Red or '2' for Black: ")
			var color int
			fmt.Scan(&color)

			if color == 1 {
				fmt.Println("You won!")
				bankroll += bet
			} else if color == 2 {
				fmt.Println("You lost.")
				bankroll -= bet
			} else {
				fmt.Println("Wrong choice")
			}
		case 2:
			bet := Bet_amount()
			fmt.Print("Enter '1' for High or '2' for Low: ")
			var Choice int
			fmt.Scan(&Choice)
			if Choice == 1 {
				fmt.Println("You won!")
				bankroll += bet
			} else if Choice == 2 {
				fmt.Println("You lost.")
				bankroll -= bet
			} else {
				fmt.Println("Invalid choice. Please try again.")
			}
		case 3:
			bet := Bet_amount()
			fmt.Print("Enter your specific number (1-36): ")
			var Number int
			fmt.Scan(&Number)
			if Number >= 1 && Number <= 36 {
				fmt.Println("You won!")
				bankroll += 35 * bet
			} else {
				fmt.Println("Invalid Choice")
			}
		case 4:
			fmt.Println("Thank you for playing!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
	fmt.Println("You're out of money")
}
