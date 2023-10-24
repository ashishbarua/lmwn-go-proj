package main

import (
	"fmt"

	"example.com/lmwn-go-proj/repository"
)

func main() {
	records := repository.GetCovidData()
	fmt.Println("1st record", records[0])
}
