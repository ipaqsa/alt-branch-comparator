package IO

import (
	"bufio"
	"os"
	"strings"
)

func inputString() string {
	print(":> ")
	msg, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.Replace(msg, "\n", "", -1)
}

func Begin() (string, string) {
	println("Hello!\nThis is CLI App for comparing packages from ALT REST API")
	println("Enter first branch`s name")
	firstName := inputString()
	println("Enter second branch` name")
	secondName := inputString()
	return firstName, secondName
}

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
