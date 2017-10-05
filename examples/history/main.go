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

	client.SessionID = auth.SessionID

	fmt.Printf("\nThe list of your velobike.ru events:\n")

	history, _, err := client.History.Get()

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		for _, item := range history.Items {
			fmt.Printf("\nType: %s\n", *item.Type)
			fmt.Printf("ID: %s\n", *item.ID)
			fmt.Printf("Date: %s\n", *item.StartDate)
			fmt.Printf("Price: %f\n", *item.Price)
			fmt.Printf("Rejected: %v\n", *item.Rejected)

			if *item.Type == "Ride" {
				fmt.Printf("BikeID: %s\n", *item.BikeID)
				fmt.Printf("BikeType: %s\n", *item.BikeType)
				fmt.Printf("Time: %s\n", *item.Time)
				fmt.Printf("Duration: %s\n", *item.Duration)
				fmt.Printf("EndDate: %s\n", *item.EndDate)
				fmt.Printf("StartBikeParkingNumber: %s\n", *item.StartBikeParkingNumber)
				fmt.Printf("StartBikeParkingName: %s\n", *item.StartBikeParkingName)
				fmt.Printf("StartBikeParkingAddress: %s\n", *item.StartBikeParkingAddress)
				fmt.Printf("StartBikeSlotNumber: %s\n", *item.StartBikeSlotNumber)
				fmt.Printf("EndBikeParkingNumber: %s\n", *item.EndBikeParkingNumber)
				fmt.Printf("EndBikeParkingName: %s\n", *item.EndBikeParkingName)
				fmt.Printf("EndBikeParkingAddress: %s\n", *item.EndBikeParkingAddress)
				fmt.Printf("EndBikeSlotNumber: %s\n", *item.EndBikeSlotNumber)
				fmt.Printf("Distance: %d meters\n", *item.Distance)
				fmt.Printf("Text: %s\n", *item.Text)
			} else if *item.Type == "Pay" {
				fmt.Printf("Contract: %s\n", *item.Contract)
				fmt.Printf("Status: %s\n", *item.Status)
				fmt.Printf("PanMask: %s\n", *item.PanMask)
			}
		}

		fmt.Printf("\nTotal Rides Time: %s\n", *history.TotalRidesTime)
	}
}
