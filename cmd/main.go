package main

import (
	"test/pkg/API"
)

func main() {
	firstSetResponse, err := API.GetSet("https://rdb.altlinux.org/api/export/branch_binary_packages/p10")
	if err != nil {
		println(err.Error())
		return
	}
	secondSetResponse, err := API.GetSet("https://rdb.altlinux.org/api/export/branch_binary_packages/p9")
	if err != nil {
		println(err.Error())
		return
	}

	err = API.Compare(firstSetResponse, secondSetResponse, "./data/difference.json")
	if err != nil {
		println(err.Error())
		return
	}
	println("Successful")
}
