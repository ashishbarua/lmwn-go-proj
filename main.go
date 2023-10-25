package main

import (
	"encoding/json"
	"fmt"

	"example.com/lmwn-go-proj/service"
)

func main() {
	cdSummary := service.GetCovidDataSummary()
	str, err := json.Marshal(cdSummary)

	if err != nil {
		fmt.Println("Marshaling failed")
	} else {
		fmt.Println(string(str))
	}
}
