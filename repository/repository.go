package repository

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"example.com/lmwn-go-proj/constants"
)

var RepositoryInstance = &Repository{}

// Some fields are set as type *any as they are consistently having "null" as their value throughout the data set
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

type IRepository interface {
	GetCovidData() []CovidRecord
}

type Repository struct{}

func (r *Repository) GetCovidData() []CovidRecord {
	apiResponse := r.callAPI()
	return r.deSerializeResponse(apiResponse)
}

func (r *Repository) callAPI() []byte {
	resp, _ := http.Get(constants.COVID_DATA_URL)
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	return body
}

func (r *Repository) deSerializeResponse(body []byte) []CovidRecord {
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
