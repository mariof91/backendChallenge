package unmarshaller

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/mariof91/backendChallenge/modules/person/model"
)

const (
	start = 9
	end   = 2
)

type PersonUnmarshaller interface {
	Unmarshal(str string) ([]model.Person, error)
}

type PersonUnmarshallerImpl struct{}

func NewPersonUnmarshallerImpl() *PersonUnmarshallerImpl {
	return &PersonUnmarshallerImpl{}
}

func (PersonUnmarshallerImpl) Unmarshal(str string) ([]model.Person, error) {
	length := len(str) - end
	s := str[start:length]
	strs := strings.Split(s, ",")

	var people []model.Person

	for i := 0; i < len(strs); i++ {
		key := strs[i]
		keyLength := len(key)
		age := strs[i+1]
		ageLength := len(age)
		aux := age[strings.Index(age, "=")+1 : ageLength]
		age2, err := strconv.Atoi(aux)
		if err != nil {
			return nil, err
		}

		person := model.Person{
			Key: key[strings.Index(key, "=")+1 : keyLength],
			Age: age2,
		}

		people = append(people, person)
		i++
	}

	return people, nil
}

type PersonUnmarshaller2Impl struct{
}

func NewPersonUnmarshaller2Impl() *PersonUnmarshaller2Impl {
	return &PersonUnmarshaller2Impl{}
}

/*
if the JSON received has the following format:

{
	"data": [{
		"key": "IAfpK",
		"age": 58
	}, {
		"key": "WNVdi",
		"age": 64
	}]
}

you should use this marshaller
*/
func (p PersonUnmarshaller2Impl) Unmarshal(str string) ([]model.Person, error) {
	people := map[string][]model.Person{}

	err:= json.Unmarshal([]byte(str), &people)
	if err != nil {
		return nil, err
	}

	return people["data"], nil
}

