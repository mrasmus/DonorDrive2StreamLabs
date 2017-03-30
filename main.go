package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/mrasmus/DonorDrive2StreamLabs/streamlabs"
	"golang.org/x/oauth2"
)

var CLIENT_ID, CLIENT_SECRET, REDIRECT_URI string

func ExampleConfig() {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		Scopes:       []string{"donations.create", "donations.read", "alerts.create"},
		Endpoint:     streamlabs.Endpoint,
		RedirectURL:  REDIRECT_URI,
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	uri := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", uri)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)
	resp, err := client.PostForm("https://streamlabs.com/api/v1.0/donations", url.Values{"name": {"mrasmus"}, "identifier": {"mrasmus@gmail.com"}, "message": {"This is a test donation"}, "amount": {"20"}, "currency": {"USD"}})
	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp, err, err1)
}

func ExampleHTTPClient() {
	hc := &http.Client{Timeout: 2 * time.Second}
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, hc)

	conf := &oauth2.Config{
		ClientID:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		Scopes:       []string{"donations.create", "donations.read", "alerts.create"},
		Endpoint:     streamlabs.Endpoint,
		RedirectURL:  REDIRECT_URI,
	}

	// Exchange request will be made by the custom
	// HTTP client, hc.
	_, err := conf.Exchange(ctx, "foo")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	CLIENT_ID = "elca3MftC4W8y9Ep83j5VsbAbYQrDygB3kNtDvxx"
	CLIENT_SECRET = "0NqViurgOEDl8DufwgZPBZL0yCAChjlr7WxBpYfJ"
	REDIRECT_URI = "http://mrasm.us/StreamLabs"

	ExampleConfig()
}
