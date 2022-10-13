package main

import (
	"encoding/json"
	"fmt"

	"github.com/chumvan/gopysense/pkg/sensehat"
)

func main() {
	m := sensehat.GetAllMeasurements()
	result, _ := json.Marshal(m)
	fmt.Println(string(result))
}
