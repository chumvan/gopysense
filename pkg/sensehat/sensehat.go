package sensehat

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// TODO Get(name string) []byte

type Measurement struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Pressure    float32 `json:"pressure"`
}

func (m Measurement) String() string {
	marshaled, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(marshaled)
}

func (m Measurement) Byte() []byte {
	marshaled, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return marshaled

}

func GetAllEnv() (m Measurement, err error) {
	out, err := exec.Command("python3", "pkg/sensehat/getAllEnvData.py").Output()
	if err != nil {
		return
	}
	fmt.Printf("From py script, received: %s\n", string(out))
	err = json.Unmarshal(out, &m)

	return
}

// func GetMeasurementOf(name string) []byte {
// 	out, err := exec.Command("python3", "pkg/sensehat/getAllEnvData.py", name).Output()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("From py script, received: %s\n", string(out))

// 	var m []byte
// 	err = json.Unmarshal(out, &m)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return m
// }
