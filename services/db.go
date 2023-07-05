package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// use my file system as a DB

func ReadStats(filename string) {

	file, err := os.Open(filename)

	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer file.Close()

	stats, err := file.Stat()

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	fmt.Printf("%v\n", stats)

}

func ReadWholeFile(fileName string) {
	content, err := ioutil.ReadFile(fileName)

	AssertNoError(err)

	fmt.Println(string(content))
}

func AssertNoError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
