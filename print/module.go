package print

import (
	"fmt"
	"Forzoom/request"
	"strings"
	"time"
)

var (
	date = string([]rune(time.Now().Format(time.RFC3339))[0:10])
	lendash = 50
)

func PrintCli(mod, focus string) (string, []Less){
	time := request.GetTime(date)
	var (
		less []Less
		str []string
	)
	switch mod{
	case "teacher":
		result, dist := request.GetTeacherSchedule(focus, date)
		if dist == 1 {
			fmt.Println("Дистант")
			for i := range result{
				moredash := make([]string, lendash-len([]rune(result[i].Lesson)))
				dash := strings.Join(moredash, "-")
				timeless := time[result[i].Num-1]
				less = append(less, Less{result[i].Zoom, timeless.Begin, timeless.End})
				str = append(str, fmt.Sprintf("%d. %s %s %s / %s\n(%s)\n", result[i].Num, result[i].Lesson, dash, timeless.Begin, timeless.End, result[i].Group))
			}
			return strings.Join(str, ""), less

		} else {
			fmt.Println("Очное")
			for i := range result{
				moredash := make([]string, lendash-len([]rune(result[i].Lesson)))
				dash := strings.Join(moredash, "-")
				timeless := time[result[i].Num-1]
				less = append(less, Less{result[i].Zoom, timeless.Begin, timeless.End})
				str = append(str, fmt.Sprintf("%d. %s %s %s / %s\n(%s, Каб.: %s)\n", result[i].Num, result[i].Lesson, dash, timeless.Begin, timeless.End, result[i].Group, result[i].Cab))
			}
			return strings.Join(str, ""), less
		}
	
	default:
		result, dist := request.GetGroupSchedule(focus, date)
		time := request.GetTime(date)
		var less []Less
		if dist == 1 {
			fmt.Println("Дистант")
			for i := range result{
				moredash := make([]string, lendash-len([]rune(result[i].Lesson)))
				dash := strings.Join(moredash, "-")
				timeless := time[result[i].Num-1]
				less = append(less, Less{result[i].Zoom, timeless.Begin, timeless.End})
				str = append(str, fmt.Sprintf("%d. %s %s %s / %s\n(%s)\n", result[i].Num, result[i].Lesson, dash, timeless.Begin, timeless.End, result[i].Teacher))
			}
			return strings.Join(str, ""), less
		
		} else {
			fmt.Println("Очное")
			for i := range result{
				moredash := make([]string, lendash-len([]rune(result[i].Lesson)))
				dash := strings.Join(moredash, "-")
				timeless := time[result[i].Num-1]
				less = append(less, Less{result[i].Zoom, timeless.Begin, timeless.End})
				str = append(str, fmt.Sprintf("%d. %s %s %s / %s\n(%s, Каб.: %s)\n", result[i].Num, result[i].Lesson, dash, timeless.Begin, timeless.End, result[i].Teacher, result[i].Cab))
			}
			return strings.Join(str, ""), less
		}
	}
}

type Less struct{
	Zoom string
	Begin string
	End string
}