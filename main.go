package main

import (
	"fmt"
	"log"

	"github.com/chumvan/gopysense/pkg/sensehat"
)

func main() {
	m, err := sensehat.GetAllMeasurements()
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%+v", m.String())
}
