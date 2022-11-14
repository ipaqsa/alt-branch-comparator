package main

import (
	"os"
	"test/pkg/API"
	"test/pkg/IO"
)

func init() {
	if !IO.Exists("./data/") {
		os.Mkdir("./data/", os.ModePerm)
	}

}
func main() {
	firstName, secondName := IO.Begin()
	firstSetResponse, err := API.GetSet(firstName)
	if err != nil {
		println(err.Error())
		return
	}
	secondSetResponse, err := API.GetSet(secondName)
	if err != nil {
		println(err.Error())
		return
	}

	templatePath := "./data/difference_"
	err = API.Compare(firstSetResponse, secondSetResponse, templatePath+firstName+"_"+secondName+".json")
	if err != nil {
		println(err.Error())
		return
	}
	println("Successful")
}
