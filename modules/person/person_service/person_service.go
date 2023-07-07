package person_service

import (
	"io/ioutil"
	"net/http"
)

const URL = "https://coderbyte.com/api/challenges/json/age-counting"

type DataService interface {
	MakeGetRequest() (string, error)
}

type DataServiceImpl struct {
}

func NewDataServiceImpl() *DataServiceImpl {
	return &DataServiceImpl{}
}

func (d *DataServiceImpl) MakeGetRequest() (string, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
