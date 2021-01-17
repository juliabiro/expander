package main

import (
	"fmt"
	"github.com/juliabiro/expander/cmd"
	"github.com/juliabiro/expander/expander"
	"os"
)

const (
	configfile = "alma.conf"
)

func main() {

	expander := expander.NewExpander(configfile)
	err := expander.ParseConfigFile()
	if err != nil {
		fmt.Printf("Failed to parse configfile %s, error is %s.", configfile, err)
	}

	args := os.Args[1:]
	input, err := cmd.ParseInput(args)

	if err != nil {
		fmt.Printf("Invalid input, %s. Error is %s.", args, err)
	}
	for _, c := range input {
		fmt.Print(expander.Expand(c))
	}
}
