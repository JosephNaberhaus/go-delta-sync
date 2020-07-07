package main

import (
	"errors"
	"flag"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/targets"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/targets/test/generate"
	"strings"
)

// A command line flag that allows an arbitrary number of key/value pairs
type FlagMap map[string]string

func (i FlagMap) Set(value string) error {
	if len(value) == 0 {
		return nil
	}

	split := strings.Split(value, ":")
	if len(split) != 2 {
		return errors.New("map flag must be in form <key>:<value>")
	}

	i[split[0]] = split[1]
	return nil
}

func (i FlagMap) String() string {
	var sb strings.Builder
	for key, value := range i {
		sb.WriteString(key)
		sb.WriteString(":")
		sb.WriteString(value)
		sb.WriteString(" ")
	}

	return sb.String()
}

func main() {
	var implementationName string
	var implementationArgs = make(FlagMap)

	flag.StringVar(&implementationName, "impl", "", "language implementation name/path to use")
	flag.Var(&implementationArgs, "implArg", "arguments to pass into implementation constructor")

	flag.Parse()

	if len(implementationName) == 0 {
		panic(errors.New("implementation name/path is required"))
	}

	implementation, err := targets.CreateImplementation(implementationName, implementationArgs)
	if err != nil {
		panic(err)
	}

	generate.GenerateAgnosticTests(implementation)
	implementation.Write("agnostic-test")

	testImplementation, err := targets.CreateTestImplementation(implementationName, implementationArgs)
	if err != nil {
		panic(err)
	}

	generate.GenerateImplementationTests(testImplementation)
	testImplementation.Write("implementation_test")
}
