package main

import (
	"github.com/golang-tire/tire/cmd"
	_ "github.com/golang-tire/tire/cmd/swagger"
)

func main() {
	cmd.Execute()
}
