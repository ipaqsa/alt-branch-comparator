package API

import (
	"encoding/json"
	"net/http"
)

var Client = &http.Client{}

func getData(url string, response interface{}) error {
	r, err := Client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(response)
}

func responseToSet(response *Response) map[Package]bool {
	SetResponse := make(map[Package]bool)
	for _, pkg := range response.Packages {
		SetResponse[pkg] = true
	}
	return SetResponse
}

func GetSet(url string) (map[Package]bool, error) {
	var response = &Response{}
	err := getData(url, response)
	if err != nil {
		return nil, err
	}
	setResponse := responseToSet(response)
	return setResponse, nil
}
