package parser

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

func directoryModuleName(directoryPath string) (moduleName string, err error) {
	if !filepath.IsAbs(directoryPath) {
		directoryPath, err = filepath.Abs(directoryPath)
		if err != nil {
			return "", err
		}
	}

	path := filepath.Join(directoryPath, "go.mod")
	if fileExists(path) {
		return goModModuleName(path)
	}

	parentPath := filepath.Dir(directoryPath)
	if parentPath == directoryPath {
		return "", errors.New("couldn't find go.mod file")
	}

	parentModule, err := directoryModuleName(parentPath)
	if err != nil {
		return "", err
	}

	return parentModule + "/" + filepath.Base(directoryPath), nil
}

func goModModuleName(goModPath string) (moduleName string, err error) {
	modFile, err := os.Open(goModPath)
	if err != nil {
		return "", err
	}

	defer modFile.Close()

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
