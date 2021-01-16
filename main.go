package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	mapping := make(map[string]string)
	data, err := ioutil.ReadFile("alma.conf")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		pairs := strings.Split(line, ":")
		mapping[pairs[0]] = strings.TrimSpace(pairs[1])
	}

	args := os.Args[1:]

	for _, c := range args {
		fmt.Print(mapping[c])
	}
}
