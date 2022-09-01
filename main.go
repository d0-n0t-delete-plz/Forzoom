package main

import (
	"fmt"
	"net/http"
	"os"
	"Forzoom/request"
)

var (
	TOKEN string = os.Getenv("TOKEN_STUDY")
	client = http.Client{}
)

func main(){
	result := request.GetGroupSchedule("9ССА-271К-21", "2022-09-02", client, TOKEN)
	for i := range result {
		fmt.Println(result[i])
	}
}