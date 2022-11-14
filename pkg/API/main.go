package API

import (
	"encoding/json"
	"errors"
	"net/http"
	"test/pkg/logger"
)

var infoLoggerAPI = logger.NewLogger("api-main", "INFO")
var errorLoggerAPI = logger.NewLogger("api-main", "ERROR")

var Client = &http.Client{}

//getData make HTTP Request to API, and decode result to Response struct.
//Takes as arguments branch`s name, arch param and Response struct
func getData(name, query string, response interface{}) error {
	template := "https://rdb.altlinux.org/api/export/branch_binary_packages/"
	r, err := Client.Get(template + name + query)
	if err != nil {
		errorLoggerAPI.Printf("getData: %s", err.Error())
		return err
	}
	if r.StatusCode != 200 {
		templateError := "branch was not found: "
		return errors.New(templateError + name + " arch: " + query)
	}
	defer r.Body.Close()
	infoLoggerAPI.Printf("Get JSON from %s", name)
	return json.NewDecoder(r.Body).Decode(response)
}

//response make map from Response struct
func responseToSet(response *Response) map[string]Package {
	SetResponse := make(map[string]Package)
	for _, pkg := range response.Packages {
		SetResponse[pkg.Name] = pkg
	}
	infoLoggerAPI.Printf("responseToSet have done")
	return SetResponse
}

// GetSet is entry point in API. Call getData and responseToSet
func GetSet(name, query string) (map[string]Package, error) {
	var response = &Response{}
	err := getData(name, query, response)
	if err != nil {
		errorLoggerAPI.Printf("GetSet %s", err.Error())
		return nil, err
	}
	setResponse := responseToSet(response)
	infoLoggerAPI.Printf("GetSet have done")
	return setResponse, nil
}
