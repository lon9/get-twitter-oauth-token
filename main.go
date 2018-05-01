package main

import (
	"bufio"
	"fmt"
	"github.com/garyburd/go-oauth/oauth"
	"os"
	"os/exec"
)

func waitInput(msg string) string {
	fmt.Print(msg)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

func main() {
	consumerKey := waitInput("Consumer Key: ")
	consumerSecret := waitInput("Consumer Secret: ")

	oauthClient := oauth.Client{
		TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
		ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
		TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
	}
	oauthClient.Credentials = oauth.Credentials{
		Token:  consumerKey,
		Secret: consumerSecret,
	}
	tempCred, err := oauthClient.RequestTemporaryCredentials(nil, "oob", nil)
	if err != nil {
		panic(err)
	}

	u := oauthClient.AuthorizationURL(tempCred, nil)
	if err := exec.Command("open", u).Run(); err != nil {
		fmt.Printf("Access  here: %s\n", u)
	}

	pin := waitInput("PIN: ")
	tokenCred, _, err := oauthClient.RequestToken(nil, tempCred, pin)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Access Token: %s\n", tokenCred.Token)
	fmt.Printf("Access Token Secret: %s\n", tokenCred.Secret)
}
