package main

import (
	"github.com/phamduytien1805/chatmodule/cmd"
	_ "github.com/phamduytien1805/chatmodule/cmd/api"
)

func main() {
	cmd.Execute()
}
