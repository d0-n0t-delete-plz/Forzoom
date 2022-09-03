package parse

import (
	"encoding/json"
	"fmt"
)

type JsonConfig struct{
	Focus string
	Mode string
	Auto bool
	Os string
}

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

type JsonTime struct{
	NumberLesson int
	Begin string
	End string
}

func ParseConfig(j []byte) JsonConfig{
	var conf []JsonConfig
	err := json.Unmarshal(j, &conf)
	if err != nil{
		fmt.Println(err)
	}
	
	return conf[0]
}

func ParseFull(j []byte) ([]JsonFull, int){
	var groups []JsonFull
	err := json.Unmarshal(j, &groups)
	if err != nil{
		fmt.Println(err)
	}
	
	return groups, groups[0].Dist
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

func ParseTime(j []byte) []JsonTime{
	var time []JsonTime
	err := json.Unmarshal(j, &time)
	if err != nil{
		fmt.Println(err)
	}

	return time
}
