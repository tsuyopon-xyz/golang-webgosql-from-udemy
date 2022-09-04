package main

import (
	"fmt"
	"log"

	"myapp/config"
)

func main() {
	fmt.Printf("%+v\n", config.Config)

	log.Println("test")
}
