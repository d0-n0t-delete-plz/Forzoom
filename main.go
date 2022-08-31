package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var TOKEN string = os.Getenv("TOKEN_STUDY")

func main(){
	client := http.Client{}

	req, err := http.NewRequest("GET", "https://api.ukrtb.ru/api/getGroups", nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("apikey", TOKEN)
	req.Header.Add("content-type", "text/html; charset=UTF-8")
	// req.Header.Add("group", "9ССА-271К-21")
	// req.Header.Add("date", "2022-09-02")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	
	buf, err := ioutil.ReadAll(resp.Body) // Need to parsing
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(strconv.Unquote(string(buf)))

}

func Parse(){} // TODO: parse from string json to slice