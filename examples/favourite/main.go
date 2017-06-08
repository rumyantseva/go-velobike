package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
	"text/tabwriter"

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

	fmt.Printf("\nThe list of your favourite parkings:\n")

	parkings, _, err := client.Parkings.List()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintf(w, "ADDRESS\tFREE BIKES\tFREE PLACES\n")
	for _, p := range parkings.Items {
		if p.IsFavourite != nil && *p.IsFavourite {
			fmt.Fprintf(w, "%s\t%d\t%d\n", *p.Address, *p.TotalPlaces-*p.FreePlaces, *p.FreePlaces)
		}
	}
	w.Flush()
}
