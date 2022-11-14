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
	println("Good, please wait for a moment")
	return firstName, secondName
}
