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
	fmt.Printf("Session ID: %s\n", *client.SessionId)
}
