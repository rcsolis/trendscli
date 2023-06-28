package internal

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// Main function for reading Google Trends data
func Run() {
	var rss RSS

	data := readGoogleTrends()
	err := xml.Unmarshal(data, &rss)
	if err != nil {
		panic(err)
	}
	fmt.Println("Google Trends found for:", rss.Print())
}

/**
 * Read Google Trends data
 * @return {[]byte} body
 */
func readGoogleTrends() []byte {
	resp := getGoogleTrends()
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}

/**
 * Make request for get Google Trends data
 * @return {http.Response} resp
 */
func getGoogleTrends() *http.Response {
	fmt.Println("Requesting for Google Trends data...")
	resp, err := http.Get(RSS_URL)
	if err != nil {
		panic(err)
	}
	return resp
}

// Path: internal/googletrends.go
