package main

import (
	"fmt"

	"github.com/rumyantseva/go-velobike/v4/velobike"
)

func main() {
	client := velobike.NewClient()

	fmt.Println("The list of the velobike.ru stations:")

	parkings, _, err := client.Parkings.List()

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		for _, item := range parkings.Items {
			fmt.Printf("%s\n", *item.Address)
		}
	}
}
