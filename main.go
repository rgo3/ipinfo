package main

import (
	"strings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	url = "http://ip-api.com/json/"
)

var commands = map[string]bool{
	"query": true,
	"city": true,
	"country": true, 
	"countryCode": true,
	"isp": true,
}

func main() {
	apiClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create request!")
		return
	}

	res, err := apiClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IP-API response! :(")
		return
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read json response!")
		return
	}

	var ipResponse map[string]interface{}
	err = json.Unmarshal([]byte(data), &ipResponse)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse json! :(")
		return
	}

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stdout, "Usage: ipinfo <cmd>")
		os.Exit(1)
	}

	cmd := os.Args[1]
	if strings.Compare(cmd, "ip") == 0 {
		cmd = "query"
	}

	if !commands[cmd] {
		fmt.Fprintf(os.Stderr, "%s is not a valid command\n", cmd)
		os.Exit(1)
	}

	fmt.Printf("%s\n", ipResponse[cmd])
	return
}
