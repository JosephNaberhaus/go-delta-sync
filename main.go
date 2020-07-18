package main

import (
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	. "github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/types"
	"github.com/JosephNaberhaus/go-delta-sync/parser"
)

func main() {
	parser.Traverse("main.go")
}

type Test struct {
	testInt int
	test    agnostic.Field
}

type Test2 struct {
	testTest Test
	test     Any
}
