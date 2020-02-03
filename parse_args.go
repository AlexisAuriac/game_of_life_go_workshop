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

var colors = map[string]string{
	"black":  "\033[1;40m",
	"red":    "\033[1;41m",
	"green":  "\033[1;42m",
	"yellow": "\033[1;43m",
	"blue":   "\033[1;44m",
	"purple": "\033[1;45m",
	"cyan":   "\033[1;46m",
	"white":  "\033[1;47m",
}

// ParsedArguments represents the parsed arguments
type ParsedArguments struct {
	speed      uint64
	colorAlive string
	colorDead  string
}

func newParsedArguments() *ParsedArguments {
	return &ParsedArguments{
		speed:      30,
		colorAlive: colors["green"],
		colorDead:  colors["red"],
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
	argument{
		name: "alive",
		parser: func(parsedArgs *ParsedArguments, arg string) error {
			if val, ok := colors[arg]; ok {
				parsedArgs.colorAlive = val
				return nil
			}

			return fmt.Errorf("%s: Invalid color", arg)
		},
	},
	argument{
		name: "dead",
		parser: func(parsedArgs *ParsedArguments, arg string) error {
			if val, ok := colors[arg]; ok {
				parsedArgs.colorDead = val
				return nil
			}

			return fmt.Errorf("%s: Invalid color", arg)
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
