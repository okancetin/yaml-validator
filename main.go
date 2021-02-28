package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("Parsing configurations")

	content, err := ioutil.ReadFile("./resources/deployment.yaml")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
