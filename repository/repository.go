package repository

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type CovidRecord struct {
	ConfirmDate    string
	No             *any
	Age            *int
	Gender         string
	GenderEn       string
	Nation         *any
	NationEn       string
	Province       string
	ProvinceId     int
	District       *any
	ProvinceEn     string
	StatQuarantine int
}

type DataResponse struct {
	Data []CovidRecord
}

const COVID_DATA_URL = "https://static.wongnai.com/devinterview/covid-cases.json"

func GetCovidData() []CovidRecord {
	apiResponse := callAPI()
	return deSerializeResponse(apiResponse)
}

func callAPI() []byte {
	resp, _ := http.Get(COVID_DATA_URL)
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	return body
}

func deSerializeResponse(body []byte) []CovidRecord {
	var dataResponse DataResponse
	var records []CovidRecord

	err := json.Unmarshal(body, &dataResponse)
	if err != nil {
		log.Fatal(err)
	} else {
		records = dataResponse.Data
	}
	return records
}
