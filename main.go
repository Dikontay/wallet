package main

import (
	"bufio"
	"fmt"
	"github.com/Dikontay/wallet/internal"
	"os"
	"strconv"
	"sync"
)

func main() {

	myWallet := &internal.Wallet{}
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nWhat would you like to do?")
		fmt.Println("1) Check Balance")
		fmt.Println("2) Deposit Bitcoin")
		fmt.Println("3) Withdraw Bitcoin")
		fmt.Println("4) Exit")
		fmt.Print("Please enter a number (1-4): ")

		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			wg.Add(1)
			go func() {
				defer wg.Done()
				fmt.Printf("Your current balance is: %.2f BTC\n", myWallet.Balance())
			}()
		case "2":
			fmt.Print("Enter the amount of Bitcoin to deposit: ")
			scanner.Scan()
			amount, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				myWallet.Deposit(internal.Bitcoin(amount))
				fmt.Printf("Deposited %.2f BTC. New balance: %.2f BTC\n", amount, myWallet.Balance())
			}()
		case "3":
			fmt.Print("Enter the amount of Bitcoin to withdraw: ")
			scanner.Scan()
			amount, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := myWallet.Withdraw(internal.Bitcoin(amount))
				if err != nil {
					fmt.Println("Unable to withdraw:", err)
				} else {
					fmt.Printf("Withdrew %.2f BTC. New balance: %.2f BTC\n", amount, myWallet.Balance())
				}
			}()
		case "4":
			fmt.Println("Exiting...")
			wg.Wait()
			return
		default:
			fmt.Println("Invalid choice, please enter a number between 1 and 4.")
		}
		wg.Wait()
	}
}
