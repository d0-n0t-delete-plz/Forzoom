package main

import (
	"Forzoom/request"
	"fmt"
	"os"
	"strings"
	// "os/exec"
)

var (
	// token = request.Token()
	token = os.Getenv("TOKEN_STUDY")
	date = "2022-09-02"
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
		if dist == 1 {
			var zoom []string
			zoom = append(zoom, result[0].Zoom)
			for i := range result{
				moredash := make([]string, lendash-len([]rune(result[i].Lesson)))
				dash := strings.Join(moredash, "-")
				fmt.Printf("%d. %s %s %s\n(%s)\n", result[i].Num, result[i].Lesson, dash, result[i].Date, result[i].Group)
			}
			return zoom
		} else {
			var zoom []string
			for i := range result{
				moredash := make([]string, lendash-len([]rune(result[i].Lesson)))
				dash := strings.Join(moredash, "-")
				fmt.Printf("%d. %s %s %s\n(%s, Каб.: %s)\n", result[i].Num, result[i].Lesson, dash, result[i].Date, result[i].Group, result[i].Cab)
			}
			return zoom
		}
	
	default:
		result, dist := request.GetGroupSchedule(focus, date, token)
		if dist == 1 {
			var zoom []string
			for i := range result{
				zoom = append(zoom, result[i].Zoom)
				moredash := make([]string, lendash-len([]rune(result[i].Lesson)))
				dash := strings.Join(moredash, "-")
				fmt.Printf("%d. %s %s %s\n(%s)\n", result[i].Num, result[i].Lesson, dash, result[i].Date, result[i].Teacher)
			}
			return zoom
		
		} else {
			var zoom []string
			for i := range result{
				moredash := make([]string, lendash-len([]rune(result[i].Lesson)))
				dash := strings.Join(moredash, "-")
				fmt.Printf("%d. %s %s %s\n(%s, Каб.: %s)\n", result[i].Num, result[i].Lesson, dash, result[i].Date, result[i].Teacher, result[i].Cab)
			}
			return zoom
		}
	}
}
