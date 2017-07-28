package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	tolstoy, _ := ioutil.ReadFile("warandpeace")
	fmt.Println(len(string(tolstoy)))
}
