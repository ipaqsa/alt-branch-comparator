package IO

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var pathToBuild = ""

func SetPathToBuild(path string) {
	pathToBuild = path
}

func GetPathToBuild() string {
	return pathToBuild
}

func inputString() string {
	print(":> ")
	msg, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	prepared := strings.Replace(msg, "\n", "", -1)
	return strings.Replace(prepared, "\r", "", -1)
}

// Begin is entry point in IO. Asks for branches`s name and data directory
func Begin() (string, string) {
	println("This is app for comparing packages from different branches ALT Linux")
	if !PathStatus {
		println("Enter path to save result(without '/' at end)")
		SetPathToBuild(inputString())
	} else {
		fmt.Printf("Result dir: %s\n", GetPathToBuild())
	}
	println("Enter first branch`s name")
	firstName := inputString()
	println("Enter second branch` name")
	secondName := inputString()
	return firstName, secondName
}

// AddQuery allows you to enter arch params
func AddQuery() string {
	println("Do you want to add platform parameters?(y/n)")
	ch := inputString()
	if ch == "y" {
		println("Enter arches separated by commas")
		arches := inputString()
		archesSplitted := strings.Split(arches, ",")
		query := "?"
		for _, arch := range archesSplitted {
			query += "arch=" + strings.TrimSpace(arch) + "&"
		}
		return query[:len(query)-1]
	} else if ch == "n" {
		return ""
	} else {
		println("Incorrect input")
		println("Parameters wont be added")
		return ""
	}
}
