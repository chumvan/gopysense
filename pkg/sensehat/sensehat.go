package sensehat

import (
	"log"
	"fmt"
	"os/exec"
)

func PrintFromPy() {
	out, err := exec.Command("python3", "pkg/sensehat/getSensorData.py").Output()	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data received: %s\n", string(out))
}