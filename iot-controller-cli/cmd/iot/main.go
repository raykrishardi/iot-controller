package main

import (
	"log"

	"github.com/raykrishardi/iot-controller-cli/internal/pkg/cmd"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cmd.Execute()
}
