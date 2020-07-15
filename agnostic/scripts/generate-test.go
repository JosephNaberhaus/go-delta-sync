package main

import (
	"errors"
	"flag"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/scripts/input"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/targets"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/test"
	"os"
)

const LimitationsFilePath = "test/limitations.json"
const LimitationsReportFilePath = "limitations.md"

// Script that takes builds both a test output file using the specified
// implementation and a file that tests the functionality of that file
func main() {
	var implementationName string
	var output string
	var testSuffix string
	var implementationArgs = make(input.Map)

	flag.StringVar(&implementationName, "impl", "", "language implementation name/path to use")
	flag.StringVar(&output, "output", "generated", "name of file to output generated test agnostic code to")
	flag.StringVar(&testSuffix, "testSuffix", "_test", "suffix to add to the end of the tests file")
	flag.Var(&implementationArgs, "implArg", "'key:value' pairs to pass as arguments to the implementation")

	flag.Parse()

	if len(implementationName) == 0 {
		panic(errors.New("implementation name/path is required"))
	}

	var suite test.Suite
	var removedCases []test.RemovedCase

	println("Looking for language limitations file")
	if _, err := os.Stat(LimitationsFilePath); err == nil {
		println("Found limitations file")
		limitations, err := test.LoadLimitations(LimitationsFilePath)
		if err != nil {
			println("Couldn't load limitations file")
			panic(err)
		}

		suite, removedCases = test.AllSuites.RemoveLimitations(limitations)
	} else {
		println("No limitations file found")

		suite = test.AllSuites
		removedCases = []test.RemovedCase{}
	}

	implementation, err := targets.CreateImplementation(implementationName, implementationArgs)
	if err != nil {
		panic(err)
	}

	println("Write test agnostic code")
	suite.GenerateAgnostic(implementation)
	implementation.Write("test/" + output)

	testImplementation, err := targets.CreateTestImplementation(implementationName, implementationArgs)
	if err != nil {
		panic(err)
	}

	println("Writing test files")
	suite.GenerateTests(testImplementation)
	testImplementation.Write("test/" + output + testSuffix)

	println("Writing Limitations Report")
	err = test.GenerateLimitationReport(LimitationsReportFilePath, removedCases)
	if err != nil {
		panic(err)
	}
}
