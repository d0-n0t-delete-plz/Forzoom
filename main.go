package main

import (
	"Forzoom/request"
	"fmt"
	"os"
	"strings"
	// "os/exec"
)

var (
	// token = request.Token() // На windows
	token = os.Getenv("TOKEN_STUDY")
	date = "2022-09-03"
	lendash = 60
)

func main(){
	conf := request.ReadConfig()
	PrintCli(conf.Mode, conf.Focus)
	fmt.Scanln()
}

func PrintCli(mod, focus string) []string{
	switch mod{
	case "teacher":
		result, dist := request.GetTeacherSchedule(focus, date, token)
		time := request.GetTime(date, token)
		if dist == 1 {
			fmt.Println("Дистант")
			var zoom []string
			zoom = append(zoom, result[0].Zoom)
			for i := range result{
				moredash := make([]string, lendash-len([]rune(result[i].Lesson)))
				dash := strings.Join(moredash, "-")
				timeless := time[result[i].Num-1]
				fmt.Printf("%d. %s %s %s / %s\n(%s)\n", result[i].Num, result[i].Lesson, dash, timeless.Begin, timeless.End, result[i].Group)
			}
			return zoom
		} else {
			fmt.Println("Очное")
			var zoom []string
			for i := range result{
				moredash := make([]string, lendash-len([]rune(result[i].Lesson)))
				dash := strings.Join(moredash, "-")
				timeless := time[result[i].Num-1]
				fmt.Printf("%d. %s %s %s / %s\n(%s, Каб.: %s)\n", result[i].Num, result[i].Lesson, dash, timeless.Begin, timeless.End, result[i].Group, result[i].Cab)
			}
			return zoom
		}
	
	default:
		result, dist := request.GetGroupSchedule(focus, date, token)
		time := request.GetTime(date, token)
		if dist == 1 {
			fmt.Println("Дистант")
			var zoom []string
			for i := range result{
				zoom = append(zoom, result[i].Zoom)
				moredash := make([]string, lendash-len([]rune(result[i].Lesson)))
				dash := strings.Join(moredash, "-")
				timeless := time[result[i].Num-1]
				fmt.Printf("%d. %s %s %s / %s\n(%s)\n", result[i].Num, result[i].Lesson, dash, timeless.Begin, timeless.End, result[i].Teacher)
			}
			return zoom
		
		} else {
			fmt.Println("Очное")
			var zoom []string
			for i := range result{
				moredash := make([]string, lendash-len([]rune(result[i].Lesson)))
				dash := strings.Join(moredash, "-")
				timeless := time[result[i].Num-1]
				fmt.Printf("%d. %s %s %s / %s\n(%s, Каб.: %s)\n", result[i].Num, result[i].Lesson, dash, timeless.Begin, timeless.End, result[i].Teacher, result[i].Cab)
			}
			return zoom
		}
	}
}
