package request

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
	"Forzoom/parse"
	"os"
)

var (
	token string = os.Getenv("TOKEN_STUDY")
	client = http.Client{}
)

func ReadConfig() parse.JsonConfig{
	file, err := os.Open("config.json")
	if err != nil {fmt.Println(err)}
	defer file.Close()

	buf := make([]byte, 250)
     
    for{
        n, err := file.Read(buf)
        if err == io.EOF{
            break
        }
		buf = buf[:n]
	}

	config := parse.ParseConfig(buf)
	return config
}

func GetGroupSchedule(group, date string) ([]parse.JsonFull, int){
	req, err := http.NewRequest("GET", "https://api.ukrtb.ru/api/getGroupSchedule?group="+group+"&date="+date, nil)
	if err != nil {fmt.Println(err)}
	req.Header.Add("apikey", token)
	
	resp, err := client.Do(req)
	if err != nil {fmt.Println(err)}
	
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {fmt.Println(err)}

	return parse.ParseFull(buf)
}

func GetTeacherSchedule(teacher, date string) ([]parse.JsonFull, int){
	req, err := http.NewRequest("GET", "https://api.ukrtb.ru/api/getTeacherSchedule?teacher="+teacher+"&date="+date, nil)
	if err != nil {fmt.Println(err)}
	req.Header.Add("apikey", token)
	
	resp, err := client.Do(req)
	if err != nil {fmt.Println(err)}
	
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {fmt.Println(err)}

	return parse.ParseFull(buf)
}

func GetGroups()[]parse.JsonGroup{
	req, err := http.NewRequest("GET", "https://api.ukrtb.ru/api/getGroups", nil)
	if err != nil {fmt.Println(err)}
	req.Header.Add("apikey", token)
	
	resp, err := client.Do(req)
	if err != nil {fmt.Println(err)}
	
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {fmt.Println(err)}

	return parse.ParseGroup(buf)
}

func GetTeachers()[]parse.JsonTeachers{
	req, err := http.NewRequest("GET", "https://api.ukrtb.ru/api/getTeachers", nil)
	if err != nil {fmt.Println(err)}
	req.Header.Add("apikey", token)
	
	resp, err := client.Do(req)
	if err != nil {fmt.Println(err)}
	
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {fmt.Println(err)}

	return parse.ParseTeacher(buf)
}

func GetTime(date string)[]parse.JsonTime{
	req, err := http.NewRequest("GET", "https://api.ukrtb.ru/api/getTime?date="+date, nil)
	if err != nil {fmt.Println(err)}
	req.Header.Add("apikey", token)
	
	resp, err := client.Do(req)
	if err != nil {fmt.Println(err)}
	
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {fmt.Println(err)}

	return parse.ParseTime(buf)
}

// func Token() string{ // В будущем удалится
// 	var Token string
// 	file, err := os.Open("token.txt")
// 	if err != nil {fmt.Println(err)}
// 	defer file.Close()

// 	buf := make([]byte, 35)
     
//     for{
//         n, err := file.Read(buf)
//         if err == io.EOF{
// 			Token = string(buf)
//             break
//         }
// 		buf = buf[:n]
// 	}
// 	return Token
// }
