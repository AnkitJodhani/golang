package main

import (
	"fmt"
)

func main() {
	// Banking application
	var balance int

	fmt.Println("Please enter the amount that you've in bank acc: ")
	fmt.Scan(&balance)

	var userChoice int
	var withdrawAmount int
	var depositAmount int
	for {
		fmt.Println("Please choose the option from blow")
		fmt.Println("Choose 1 for checking balance: ")
		fmt.Println("Choose 2 for withdrawing: ")
		fmt.Println("Choose 3 for deposit : ")
		fmt.Println("Choose 4 to exit: ")
		fmt.Scan(&userChoice)

		switch userChoice {
		case 1:
			fmt.Println("Your Account balancer is:", balance)
		case 2:
			fmt.Println("Eneter the amount that you want to withdraw:")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > balance {
				fmt.Println("Eneter valid amount")
			} else {
				balance = balance - withdrawAmount
				fmt.Println("Take your cache $", withdrawAmount)
				fmt.Println("Current balance is ", balance)
			}
		case 3:
			fmt.Println("Please enter the amount that you want to deposit: ")
			fmt.Scan(&depositAmount)
			fmt.Println("Please enter your cache: ")
			balance = balance + depositAmount
			fmt.Println("Deposit sucessful 🙌")
			fmt.Println("Current balancer is ", balance)

		case 4:
			fmt.Println("Thank you for using our ATM")
		}
		if userChoice == 4 {
			break
		}
	}
}
