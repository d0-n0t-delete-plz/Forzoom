package request

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"Forzoom/parse"
	"os"
)

var (
	token string = os.Getenv("TOKEN_STUDY")
	client = http.Client{}
)

func GetGroupSchedule(group string, date string) []parse.JsonFull{
	
	req, err := http.NewRequest("GET", "https://api.ukrtb.ru/api/getGroupSchedule?group="+group+"&date="+date, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("apikey", token)
	
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

func GetTeacherSchedule(teacher string, date string) []parse.JsonFull{
	req, err := http.NewRequest("GET", "https://api.ukrtb.ru/api/getTeacherSchedule?teacher="+teacher+"&date="+date, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("apikey", token)
	
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(buf))

	return parse.ParseFull(buf)
}

func GetGroups()[]parse.JsonGroup{
	req, err := http.NewRequest("GET", "https://api.ukrtb.ru/api/getGroups", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("apikey", token)
	
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return parse.ParseGroup(buf)
}

func GetTeachers()[]parse.JsonTeachers{
	req, err := http.NewRequest("GET", "https://api.ukrtb.ru/api/getTeachers", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("apikey", token)
	
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return parse.ParseTeacher(buf)
}