package parser

import (
	"bufio"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func Traverse(fileName string) error {
	fSet := token.NewFileSet()
	file, err := parser.ParseFile(fSet, fileName, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	moduleName, err := getModuleName(filepath.Dir(fileName))
	if err != nil {
		return err
	}

	typeSpecs := make(map[Type]*ast.TypeSpec)

	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					typeName := typeSpec.Name.Name
					if _, ok = knownType[typeName]; ok {
						return errors.New("duplicate type name \"" + typeName + "\"")
					}

					knownType[typeName] = true

					//switch typeStruct := typeSpec.Type.(type) {
					//case *ast.StructType:
					//	for
					//}
				}
			}
		}
	}

	println("done parsing")
	return nil
}

func getModuleName(directoryName string) (moduleName string, err error) {
	modFile, err := os.Open(filepath.Join(directoryName, "go.mod"))
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(modFile)
	for scanner.Scan() {
		lineTokens := strings.Fields(scanner.Text())
		if len(lineTokens) > 0 && lineTokens[0] == "module" {
			if len(lineTokens) > 1 {
				return lineTokens[1], nil
			} else {
				return "", errors.New("no module name found after module verb")
			}
		}
	}

	return "", errors.New("no module verb found")
}

func findTypeSpecs(file *ast.File, typeSpecs)