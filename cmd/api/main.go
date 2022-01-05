package main

import (
	"log"

	"github.com/DanielQuerolBeltran/Climbing-notebook-api/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}