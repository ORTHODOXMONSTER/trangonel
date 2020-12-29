package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("banner.txt")
	if err != nil {
		log.Fatalln("No banner file found.")
		panic(err)
	}

	fmt.Println(string(data))
}
