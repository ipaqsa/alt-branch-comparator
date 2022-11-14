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
	//Get names of the branches
	firstName, secondName := IO.Begin()

	//Get arch params
	query := IO.AddQuery()

	println("Good, please wait for a moment")
	//Make request to first branch
	firstSetResponse, err := API.GetSet(firstName, query)
	if err != nil {
		println(err.Error())
		return
	}

	//Make request to second branch
	secondSetResponse, err := API.GetSet(secondName, query)
	if err != nil {
		println(err.Error())
		return
	}

	//Compare, save and print results
	templatePath := "./data/difference_"
	err = API.Compare(firstSetResponse, secondSetResponse, templatePath+firstName+"_"+secondName+".json")
	if err != nil {
		println(err.Error())
		return
	}
	println("Successful")
}
