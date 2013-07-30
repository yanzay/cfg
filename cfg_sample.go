package main

import (
  "cfg"
	"fmt"
	"log"
)

func main() {
	mymap := make(map[string]string)
	err := cfg.Load("test.cfg", mymap)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", mymap)
}
