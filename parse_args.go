// An implementation of Conway's Game of Life.
package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ParsedArguments represents the parsed arguments
type ParsedArguments struct {
	speed uint64
}

func newParsedArguments() *ParsedArguments {
	return &ParsedArguments{
		speed: 30,
	}
}

type argumentParser func(*ParsedArguments, string) error

type argument struct {
	name   string
	parser argumentParser
}

var arguments = [...]argument{
	argument{
		name: "speed",
		parser: func(parsedArgs *ParsedArguments, arg string) error {
			speed, err := strconv.ParseUint(arg, 10, 16)

			if speed > 300 {
				speed = 300
			} else if speed == 0 {
				return errors.New("Speed can't be null")
			}

			parsedArgs.speed = speed
			return err
		},
	},
}

// ParseArgs parses the arguments and returns a structure representing the result
func ParseArgs() (*ParsedArguments, error) {
	parsedArgs := newParsedArguments()
	args := os.Args[1:]

	for _, arg := range args {
		for _, i := range arguments {
			match := fmt.Sprintf("^--%s=", i.name)
			matched, _ := regexp.MatchString(match, args[0])

			if !matched {
				continue
			}

			split := strings.SplitN(arg, "=", 2)

			if len(split) != 2 {
				return nil, fmt.Errorf("No argument specified for option '%s'", i.name)
			}

			err := i.parser(parsedArgs, split[1])

			if err != nil {
				return nil, err
			}

			break
		}
	}

	return parsedArgs, nil
}
