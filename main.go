package main

import (
	"adventlib/inputfetcher",
	"fmt"
)

func main() {

	bearerToken := "your-bearer-token-here"
	fetcher := inputfetcher.NewDataFetcher(bearerToken)

	data, err := fetcher.FetchDay1Data()

	if err != nil {
		err := fmt.Errorf("Error Fetching data: %w", err)
		fmt.Println(err)
		return
	}

	fmt.Printf("Data: %s", data)

}
