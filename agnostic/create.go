package agnostic

import (
	"errors"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/targets/golang"
)

func CreateImplementation(name string, args map[string]string) (implementation blocks.Implementation, err error) {
	if name == "go" {
		return golang.Implementation(args), nil
	}

	return nil, errors.New("No implementation found for \"" + name + "\"")
}
