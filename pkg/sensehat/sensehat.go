package sensehat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"os/exec"
)

// TODO Get(name string) []byte

type Measurement struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Pressure    float32 `json:"pressure"`
}

type Orientation struct {
	Pitch float64 `json:"pitch"`
	Roll  float64 `json:"roll"`
	Yaw   float64 `json:"yaw"`
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

func (m Orientation) String() string {
	marshaled, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(marshaled)
}

func (m Orientation) Byte() []byte {
	marshaled, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return marshaled

}

func (m Orientation) LowerRes() (lower Orientation) {
	lower = m
	lower.Pitch = float64(math.Round(m.Pitch))
	lower.Roll = float64(math.Round(m.Roll))
	lower.Yaw = float64(math.Round(m.Yaw))
	return lower

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

func GetOrientation() (m Orientation, err error) {
	out, err := exec.Command("python3", "pkg/sensehat/getOrientation.py").Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("From py script, received: %s\n", string(out))
	err = json.Unmarshal(out, &m)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func GetOrientationDebug() (m Orientation, err error) {
	cmd := exec.Command("python3", "pkg/sensehat/getOrientation.py")

	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	err = cmd.Run()
	if err != nil {
		panic(fmt.Sprint(err) + ":" + stdErr.String())
	}
	out := stdOut.Bytes()
	fmt.Printf("From py script, received: %s\n", string(out))

	err = json.Unmarshal(out, &m)
	if err != nil {
		panic(err)
	}
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
