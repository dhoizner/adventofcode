package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day02.input", "Relative path to input file.")

type Command struct {
	command string
	input   int
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		panic(err)
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	commands := make([]Command, len(lines))
	for i, v := range lines {
		split := strings.Split(v, " ")
		command := split[0]
		input, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		commands[i] = Command{command: command, input: input}
	}

	xPos, depth := 0, 0
	for _, c := range commands {
		switch c.command {
		case "forward":
			xPos += c.input
		case "down":
			depth += c.input
		case "up":
			depth -= c.input
		}
	}

	fmt.Println("xPos:", xPos, "yPos:", depth, "position:", xPos*depth)

	xPos, depth, aim := 0, 0, 0
	for _, c := range commands {
		switch c.command {
		case "forward":
			xPos += c.input
			depth += aim * c.input
		case "down":
			aim += c.input
		case "up":
			aim -= c.input
		}
	}

	fmt.Println("xPos:", xPos, "yPos:", depth, "aim:", aim, "position:", xPos*depth)
}
