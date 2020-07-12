package main

import (
	"errors"
	"flag"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/scripts/input"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/targets"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/test"
)

// Script that takes builds both a test output file using the specified
// implementation and a file that tests the functionality of that file
func main() {
	var implementationName string
	var implementationArgs = make(input.Map)

	flag.StringVar(&implementationName, "impl", "", "language implementation name/path to use")
	flag.Var(&implementationArgs, "implArg", "'key:value' pairs to pass as arguments to the implementation")

	flag.Parse()

	if len(implementationName) == 0 {
		panic(errors.New("implementation name/path is required"))
	}

	implementation, err := targets.CreateImplementation(implementationName, implementationArgs)
	if err != nil {
		panic(err)
	}

	test.AllSuites.GenerateAgnostic(implementation)
	implementation.Write("test/agnostic-test")

	testImplementation, err := targets.CreateTestImplementation(implementationName, implementationArgs)
	if err != nil {
		panic(err)
	}

	test.AllSuites.GenerateTests(testImplementation)
	testImplementation.Write("test/implementation_test")
}
