package targets

import (
	"errors"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/targets/golang"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/test"
)

func CreateImplementation(name string, args map[string]string) (implementation agnostic.Implementation, err error) {
	if name == "go" {
		return golang.NewImplementation(args), nil
	}

	return nil, errors.New("No implementation found for \"" + name + "\"")
}

func CreateTestImplementation(name string, args map[string]string) (implementation test.Implementation, err error) {
	if name == "go" {
		return golang.TestImplementation(args), nil
	}

	return nil, errors.New("No test implementation found for \"" + name + "\"")
}
