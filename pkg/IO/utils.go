package IO

import (
	"flag"
	"fmt"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

var PathStatus = false
var Debug = false

func ArgsParse() {
	debug := flag.Bool("d", false, "debug info")
	path := flag.String("p", "", "Enter path to result-dir(without / at end)")
	flag.Parse()
	fmt.Printf("path: %s\n", *path)
	fmt.Printf("debug: %t\n", *debug)
	Debug = *debug
	if *path != "" {
		if Exists(*path) {
			PathStatus = true
			SetPathToBuild(*path)
		}
	}

}
