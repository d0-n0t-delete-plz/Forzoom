package main

import (
	"Forzoom/request"
	"Forzoom/print"
	"fmt"
	// "os/exec"
)

var delay = 5

func main(){
	conf := request.ReadConfig()
	str, less := print.PrintCli(conf.Mode, conf.Focus)

	fmt.Println(str)
	fmt.Println(less)

	fmt.Scanln()
}

func SchedulerLess(str string, less []print.Less){}
