package service

import (
	"fmt"
	"reflect"
	"testing"

	"example.com/lmwn-go-proj/repository"
)

func TestGetCovidDataSummary(t *testing.T) {
	testservice := Service{
		Repository: MockRepoInstance,
	}

	summary := testservice.GetCovidDataSummary()
	fmt.Println("Summary", summary)

	expectedsummary := CountResponse{
		Province: ProvinceType{
			"A": 1,
			"B": 2,
			"C": 2,
			"D": 1,
		},
		AgeGroup: AgeGroupType{
			ZeroToThirty:     1,
			ThirtyOneToSixty: 4,
			SixtyOnePlus:     1,
			NA:               0,
		},
	}

	if !reflect.DeepEqual(summary, expectedsummary) {
		t.Errorf("Test failed: Expected %v, received %v", expectedsummary, summary)
	} else {
		t.Logf("Test passed: Expected %v, received %v", expectedsummary, summary)
	}
}

func TestUpdateAgeGroupTally(t *testing.T) {
	age1 := 19
	age2 := 45
	age3 := 65
	age4 := 0
	c1 := repository.CovidRecord{Age: &age1}
	c2 := repository.CovidRecord{Age: &age2}
	c3 := repository.CovidRecord{Age: &age3}
	c4 := repository.CovidRecord{Age: &age4}
	c5 := repository.CovidRecord{}
	c6 := repository.CovidRecord{Age: &age2}
	c7 := repository.CovidRecord{Age: &age3}
	c8 := repository.CovidRecord{Age: &age4}

	recordCollections := []repository.CovidRecord{c1, c2, c3, c4, c5, c6, c7, c8}
	agegrouptally := AgeGroupType{}
	for _, rc := range recordCollections {
		agegrouptally.updateAgeGroupTally(rc)
	}

	expectedTally := AgeGroupType{
		ZeroToThirty:     3,
		ThirtyOneToSixty: 2,
		SixtyOnePlus:     2,
		NA:               1,
	}

	if !reflect.DeepEqual(expectedTally, agegrouptally) {
		t.Errorf("Test failed: Expected %v, received %v", expectedTally, agegrouptally)
	} else {
		t.Logf("Test passed: Expected %v, received %v", expectedTally, agegrouptally)
	}
}

func TestUpdateProvinceTally(t *testing.T) {
	c1 := repository.CovidRecord{Province: "A"}
	c2 := repository.CovidRecord{Province: "B"}
	c3 := repository.CovidRecord{Province: "C"}
	c4 := repository.CovidRecord{Province: "A"}
	c5 := repository.CovidRecord{Province: "B"}
	c6 := repository.CovidRecord{Province: "D"}

	recordCollections := []repository.CovidRecord{c1, c2, c3, c4, c5, c6}
	provincetally := ProvinceType{}
	for _, rc := range recordCollections {
		provincetally.updateProvinceTally(rc)
	}

	expectedTally := ProvinceType{
		"A": 2,
		"B": 2,
		"C": 1,
		"D": 1,
	}

	if !reflect.DeepEqual(expectedTally, provincetally) {
		t.Errorf("Test failed: Expected %v, received %v", expectedTally, provincetally)
	} else {
		t.Logf("Test passed: Expected %v, received %v", expectedTally, provincetally)
	}
}
