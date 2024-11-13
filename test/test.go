package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Response struct {
	Data map[string]interface{} `json:"data"`
}

func main() {
	// Make an HTTP request
	resp, err := http.Get("https://foooooo.free.beeceptor.com")

	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	jsonstr := string(body)
	cleanedData := strings.ReplaceAll(jsonstr, "\x00", "")

	// Print the raw response body
	fmt.Println("Raw response body:", jsonstr)
	fmt.Println("XRaw response body:", cleanedData)
	// Unmarshal the JSON response
	var jsonResp Response

	err = json.Unmarshal(body, &jsonResp)

	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Access the data
	fmt.Println("Data:", jsonResp.Data)

}
