package service

import "example.com/lmwn-go-proj/repository"

type MockRepo struct{}

func (m *MockRepo) GetCovidData() []repository.CovidRecord {
	age23 := 23
	age33 := 33
	age43 := 43
	age63 := 63
	return []repository.CovidRecord{
		{Age: &age23, Province: "A"},
		{Age: &age33, Province: "B"},
		{Age: &age43, Province: "C"},
		{Age: &age63, Province: "D"},
		{Age: &age33, Province: "B"},
		{Age: &age43, Province: "C"},
	}
}

var MockRepoInstance = &MockRepo{}
