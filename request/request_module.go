package request

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"Forzoom/parse"
)

func GetGroupSchedule(group string, date string, client http.Client, TOKEN string) []parse.JsonFull{
	
	req, err := http.NewRequest("GET", "https://api.ukrtb.ru/api/getGroupSchedule?group="+group+"&date="+date, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("apikey", TOKEN)
	
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return parse.ParseFull(buf)
}