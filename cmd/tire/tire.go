package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/golang-tire/tire/internal/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	command := cmd.NewTireCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
