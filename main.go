package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func getBearerToken() (string, error) {
	// os.Args[0] is the program name, and subsequent elements are the arguments
	if len(os.Args) < 2 {
		return "", errors.New("Need to pass the bearer token in as arg 1")
	}
	return os.Args[1], nil
}

func main() {

	bearerToken, err := getBearerToken()
	if err != nil {
		fmt.Println("Error getting bearer token", err)
		os.Exit(-1)
	}

	fetcher := inputfetcher.NewDataFetcher(bearerToken)

	data, err := fetcher.FetchDay1Data()

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		return
	}

	fmt.Printf("Data: %s", data)

}
