package sensehat

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

type Measurement struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Pressure    float32 `json:"pressure"`
}

func GetMeasurements() Measurement {
	out, err := exec.Command("python3", "pkg/sensehat/getSensorData.py").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("From py script, received: %s\n", string(out))

	var m Measurement
	err = json.Unmarshal(out, &m)
	if err != nil {
		log.Fatal(err)
	}

	return m
}
