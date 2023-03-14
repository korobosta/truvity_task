package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

type UrlResponse struct {
    url string
    size int
}

func processUrls(urls []string) []UrlResponse{
    // Create urlResponse to hold single response
    var urlResponse UrlResponse
    // Create urlResponses to hold all responses
    var urlResponses []UrlResponse
    var url string;

    //Iterate the urls to make get request for each
    for e := 0; e < len(urls); e++ {
        url = urls[e]
        resp, err := http.Get(url)
        if err != nil {
            log.Fatalln(err)
        }
        //We Read the response body on the line below.
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatalln(err)
        }
        //Find size of a string
        size := len(body) 
        //Create urlResponse object
        urlResponse.size = size
        urlResponse.url = url
        //Append object to existing urlResponses
        urlResponses = append(urlResponses, urlResponse)
    }

    //sort
    sortResponses(urlResponses)
    //Return the list of responses
    return urlResponses
}

func sortResponses (urlResponses []UrlResponse) []UrlResponse{
    //sort bases on size
    sort.SliceStable(urlResponses, func(i, j int) bool {
		return urlResponses[i].size < urlResponses[j].size
	})

    return urlResponses
}

func main() {

    //Get input from user
    fmt.Println("Enter urls separated by comma:")
    var text string

    _, err := fmt.Scan(&text)
    if err != nil {
        log.Fatal(err)
    }
    data := strings.Split(text, ",")
    //Call the processUrls function the make the get requests
    response := processUrls(data)
    //Print the response
    log.Println(response)
}