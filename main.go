package main

import (
	"fmt"

	"github.com/chumvan/gopysense/pkg/sensehat"
)

func main() {
	fmt.Println("Testing py script in go exec")
	sensehat.PrintFromPy()
}
