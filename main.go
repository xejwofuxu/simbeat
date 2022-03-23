package main

import (
	"os"

	"github.com/xejwofuxu/simbeat/cmd"

	_ "github.com/xejwofuxu/simbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
