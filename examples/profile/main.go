package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/rumyantseva/go-velobike/v3/velobike"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	// First of all we need to authorize user
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

	// If we set client.SessionID we can use "only authorized" methods:
	profile, _, err := client.Profile.Get()

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		fmt.Printf("\nUser ID: %s\n", *profile.UserID)
		fmt.Printf("Email: %s\n", *profile.Email)
	}
}
