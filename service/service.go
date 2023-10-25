package service

import (
	"example.com/lmwn-go-proj/repository"
)

type IService interface {
	GetCovidDataSummary() CountResponse
}

type Service struct {
	Repository repository.IRepository
}

func (s *Service) GetCovidDataSummary() CountResponse {
	records := s.Repository.GetCovidData()
	var agTally AgeGroupType
	provinceTally := ProvinceType{}

	for _, r := range records {
		agTally.updateAgeGroupTally(r)
		provinceTally.updateProvinceTally(r)
	}

	return CountResponse{
		AgeGroup: agTally,
		Province: provinceTally,
	}
}

type AgeGroupType struct {
	ZeroToThirty     int `json:"0-30"`
	ThirtyOneToSixty int `json:"31-60"`
	SixtyOnePlus     int `json:"61+"`
	NA               int `json:"N/A"`
}

func (ag *AgeGroupType) updateAgeGroupTally(cr repository.CovidRecord) {
	if cr.Age == nil {
		ag.NA += 1
	} else if *cr.Age < 31 {
		ag.ZeroToThirty += 1
	} else if *cr.Age < 61 {
		ag.ThirtyOneToSixty += 1
	} else {
		ag.SixtyOnePlus += 1
	}
}

type ProvinceType map[string]int

func (p ProvinceType) updateProvinceTally(cr repository.CovidRecord) {
	p[cr.Province] += 1
}

type CountResponse struct {
	AgeGroup AgeGroupType
	Province ProvinceType
}

var ServiceInstance = &Service{
	Repository: repository.RepositoryInstance,
}
