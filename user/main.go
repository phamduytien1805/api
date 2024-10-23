package main

import (
	"github.com/phamduytien1805/usermodule/cmd"

	_ "github.com/phamduytien1805/usermodule/cmd/api"
)

func main() {
	cmd.Execute()
}
