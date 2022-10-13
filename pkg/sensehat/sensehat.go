package sensehat

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

func GetAllMeasurements() []byte {
	out, err := exec.Command("python3", "pkg/sensehat/getAllEnvData.py").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("From py script, received: %s\n", string(out))

	var m []byte
	err = json.Unmarshal(out, &m)
	if err != nil {
		log.Fatal(err)
	}

	return m
}

func GetMeasurementOf(name string) []byte {
	out, err := exec.Command("python3", "pkg/sensehat/getAllEnvData.py").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("From py script, received: %s\n", string(out))

	var m []byte
	err = json.Unmarshal(out, &m)
	if err != nil {
		log.Fatal(err)
	}

	return m
}
