package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/rumyantseva/go-velobike/velobike"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Velobike.ru User ID: ")
	userid, _ := r.ReadString('\n')

	fmt.Print("Velobike.ru Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	tp := velobike.BasicAuthTransport{
		Username: strings.TrimSpace(userid),
		Password: strings.TrimSpace(password),
	}

	client := velobike.NewClient(tp.Client())
	auth, _, err := client.Authorization.Authorize()

	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	client.SessionId = auth.SessionId

	fmt.Printf("\nThe list of your velobike.ru events:\n")

	history, _, err := client.History.Get()

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		for _, item := range history.Items {
			fmt.Printf("\nType: %s\n", *item.Type)
			fmt.Printf("Date: %s\n", *item.StartDate)
			fmt.Printf("Price: %f\n", *item.Price)

			if *item.Type == "Ride" {
				fmt.Printf("Time: %s/Duration: %s\n", *item.Time, *item.Duration)
				fmt.Printf("EndDate: %s\n", *item.EndDate)
				fmt.Printf("StartBikeParkingNumber: %s\n", *item.StartBikeParkingNumber)
				fmt.Printf("EndBikeParkingNumber: %s\n", *item.EndBikeParkingNumber)
				fmt.Printf("BikeID: %s\n", *item.BikeID)
				fmt.Printf("Distance: %d meters\n", *item.Distance)
			}
		}

		fmt.Printf("Total Rides Time: %s\n", *history.TotalRidesTime)
	}
}
