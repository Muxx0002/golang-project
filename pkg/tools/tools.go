package tools

import (
	"log"
	"os"
)

func CreateLogFile() *os.File {
	file, err := os.OpenFile("LOG.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
