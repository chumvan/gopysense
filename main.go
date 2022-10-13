package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/chumvan/gopysense/pkg/sensehat"
)

type Measurement struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Pressure    float32 `json:"pressure"`
}

func main() {
	out := sensehat.GetAllMeasurements()
	var result Measurement
	err := json.Unmarshal(out, &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", result)
	fmt.Println(string(out))
}
