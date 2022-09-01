package main

import (
	"fmt"
	"Forzoom/request"
)

func main(){
	test2()
}

func test1(){
	groupschedule := request.GetGroupSchedule("9ССА-271К-21", "2022-09-03")
	for i := range groupschedule {
		fmt.Println(groupschedule[i])
	}
}

func test2(){
	teacherschedule := request.GetTeacherSchedule("Баймурзина Минзаля Зайнулловна", "2022-09-02")
	for i := range teacherschedule {
		fmt.Println(teacherschedule[i])
	}
}

func test3(){
	groups := request.GetGroups()
	for i := range groups {
		fmt.Println(groups[i])
	}
}

func test4(){
	teachers := request.GetTeachers()
	for i := range teachers {
		fmt.Println(teachers[i])
	}
}