package main

import (
	"fmt"

	"github.com/chumvan/gopysense/pkg/sensehat"
)

func main() {
	m, err := sensehat.GetAllEnv()
	if err != nil {
		panic(err)
	}
	fmt.Println(m.String())
}
