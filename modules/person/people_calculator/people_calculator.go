package people_calculator

import "github.com/mariof91/backendChallenge/modules/person/model"

type PeopleCalculator interface {
	GetQuantityOfPeopleOlderThanFifty(people []model.Person) int
}

type PeopleCalculatorImpl struct{}

func NewPeopleCalculatorImpl() *PeopleCalculatorImpl {
	return &PeopleCalculatorImpl{}
}

func (PeopleCalculatorImpl) GetQuantityOfPeopleOlderThanFifty(people []model.Person) int {
	counter := 0
	for _,person := range people{
		if person.Age > 50 {
			counter ++
		}
	}
	return counter
}

