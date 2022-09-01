package parse

import (
	"encoding/json"
	"fmt"
)

type JsonFull struct{
	Num int
	Date string
	Group string
	Lesson string
	Teacher string
	Cab string
	Dist int
	Zoom string
}

type JsonGroup struct{
	Group string
}

type JsonTeachers struct{
	Fio string
	Zoom string
}

func ParseFull(j []byte) []JsonFull{
	var groups []JsonFull
	err := json.Unmarshal(j, &groups)
	if err != nil{
		fmt.Println(err)
	}
	
	return groups
}

func ParseGroup(j []byte) []JsonGroup{
	var groups []JsonGroup
	err := json.Unmarshal(j, &groups)
	if err != nil{
		fmt.Println(err)
	}

	return groups
}

func ParseTeacher(j []byte) []JsonTeachers{
	var teachers []JsonTeachers
	err := json.Unmarshal(j, &teachers)
	if err != nil{
		fmt.Println(err)
	}

	return teachers
}