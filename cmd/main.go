package main

import (
	"altComparator/pkg"
	"altComparator/pkg/API"
	"altComparator/pkg/IO"
	"fmt"
)

func init() {
	pkg.SetVersion("0.7.0")
	IO.PathFromArg()
}

func main() {
	fmt.Printf("ALT tool V-%s\n", pkg.GetVersion())
	//Get names of the branches

	firstName, secondName := IO.Begin()

	//Get arch params
	query := IO.AddQuery()

	println("Good, please wait a couple of minutes\n")
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
	templatePath := IO.GetPathToBuild() + "/difference_"
	err = API.Compare(firstName, secondName, firstSetResponse, secondSetResponse, templatePath+firstName+"_"+secondName+".json")
	if err != nil {
		println(err.Error())
		return
	}
	println("Successful")
}
