package main

import (
	"fmt"

	"github.com/mariof91/backendChallenge/modules/person/people_calculator"
	"github.com/mariof91/backendChallenge/modules/person/person_service"
	"github.com/mariof91/backendChallenge/modules/person/unmarshaller"
)

func main() {
	peopleService := person_service.NewDataServiceImpl()
	unmarshaller := unmarshaller.NewPersonUnmarshallerImpl()
	peopleCalculator := people_calculator.NewPeopleCalculatorImpl()


	rawData, err := peopleService.MakeGetRequest()
	if err != nil  {
		panic(err)
	}

	people, err := unmarshaller.Unmarshal(rawData)
	if err != nil  {
		panic(err)
	}

	quantity := peopleCalculator.GetQuantityOfPeopleOlderThanFifty(people)

	fmt.Println(quantity)
}
