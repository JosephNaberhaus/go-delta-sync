package input

import (
	"errors"
	"strings"
)

// A command line input that allows an arbitrary number of key/value pairs
type Map map[string]string

func (i Map) Set(value string) error {
	if len(value) == 0 {
		return nil
	}

	split := strings.Split(value, ":")
	if len(split) != 2 {
		return errors.New("map input must be in form <key>:<value>")
	}

	i[split[0]] = split[1]
	return nil
}

func (i Map) String() string {
	var sb strings.Builder
	for key, value := range i {
		sb.WriteString(key)
		sb.WriteString(":")
		sb.WriteString(value)
		sb.WriteString(" ")
	}

	return sb.String()
}
