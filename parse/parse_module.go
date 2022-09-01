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
	Dist bool
	Zoom string
}

type JsonTime struct{
	Num int
	Start string
	End string
}

type JsonGroups struct{
	Group string
}

func ParseFull(j []byte) []JsonFull{
	var groups []JsonFull
	err := json.Unmarshal(j, &groups)
	if err != nil{
		fmt.Println(err)
	}

	return groups
}