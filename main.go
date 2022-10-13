package main

import (
	"github.com/chumvan/gopysense/pkg/sensehat"
	"encoding/json"
	"fmt"
)

func main() {
	m := sensehat.GetMeasurements()
	result, _ := json.Marshal(m)
	fmt.Println(string(result))
}
