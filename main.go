package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bitrise-io/bitrise-oauth/client"
)

func main() {
	fmt.Print("Hello from my auth poc!")
	var realmOption client.Option = client.WithRealm("anka")
	baseUrlOption := client.WithBaseURL("http://192.168.1.4:8080")
	scope := client.WithScope("groups")
	authProvider := client.NewWithSecret(
		"anka-controller",
		"b0bbc6cf-8fdf-4f27-81f9-5bf99453b704",
		scope,
		baseUrlOption,
		realmOption)

	// tokenSource := authProvider.UMATokenSource()
	tokenSource := authProvider.TokenSource()
	token, err := tokenSource.Token()
	// token, err := tokenSource.Token(
	// 	nil,
	// 	make([]client.Permission, 0),
	// 	config.NewAudienceConfig("anka-controller"))

	if err != nil {
		fmt.Println()
		fmt.Printf("Could not get a token! Error: %q\n", err)
		os.Exit(1)
	}
	fmt.Println()
	fmt.Println()
	fmt.Printf("AccessToken: %s\n", token.AccessToken)
	httpClient := authProvider.HTTPClient()
	//httpClient := &http.Client{}

	req, _ := http.NewRequest("GET", "http://192.168.1.110/api/v1/node", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("\n\n\nError happened during the request: %s", err)
	}
	//fmt.Printf("\n\n\nRequest: %+v.\n\nHeaders: %+v, (Count: %d)",
	//	req,
	//	req.Header,
	//	len(req.Header))
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	fmt.Printf("\n\n\nResponse: %+v. \n\nBody: %q", resp, bodyString)
}
