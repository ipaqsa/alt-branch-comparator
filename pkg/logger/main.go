package logger

import (
	"altComparator/pkg/IO"
	"fmt"
	"log"
	"os"
	"time"
)

func NewLogger(object, typeLogger string) *log.Logger {
	if !IO.Exists("/var/log/altbranchcompare.d") {
		err := os.Mkdir("/var/log/altbranchcompare.d", os.ModePerm)
		if err != nil {
			println(err.Error())
			os.Exit(-1)
		}
	}
	logsDir := "/var/log/altbranchcompare.d/" + object + "/"
	if IO.Exists(logsDir) == false {
		err := os.Mkdir(logsDir, os.ModePerm)
		if err != nil {
			println(err.Error())
			os.Exit(-1)
		}
	}
	year, month, day := time.Now().Date()
	filename := fmt.Sprintf("%v-%v-%v.log", day, month.String(), year)
	path := logsDir + filename

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		println("ERROR: Create|Open File", path)
		os.Exit(-1)
	}
	if typeLogger == "ERROR" {
		return log.New(f, typeLogger+": ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		return log.New(f, typeLogger+": ", log.Ldate|log.Ltime)
	}
}
