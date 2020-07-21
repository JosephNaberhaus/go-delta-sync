package parser

import (
	"bufio"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

const tempDirPattern = "go-delta-sync_test"
const testModuleName = "test.com/module"

func createTestGoMod(directoryPath string) error {
	file, err := os.Create(filepath.Join(directoryPath, "go.mod"))
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("module " + testModuleName)
	if err != nil {
		file.Close()
		return err
	}

	err = writer.Flush()
	if err != nil {
		file.Close()
		return err
	}

	return file.Close()
}

func TestFileExists(t *testing.T) {
	directoryName, err := ioutil.TempDir("", tempDirPattern)
	require.NoError(t, err)

	defer func() {
		err := os.RemoveAll(directoryName)
		require.NoError(t, err)
	}()

	file, err := os.Create(filepath.Join(directoryName, tempDirPattern))
	require.NoError(t, err)
	file.Close()

	require.Equal(t, true, fileExists(file.Name()))
	require.Equal(t, false, fileExists(filepath.Join(directoryName, "nonexistent")))
}

func TestDirectoryModuleName(t *testing.T) {
	parentDirectory, err := ioutil.TempDir("", tempDirPattern)
	require.NoError(t, err)
	childDirectory := filepath.Join(parentDirectory, "child")
	err = os.Mkdir(childDirectory, os.ModePerm)
	require.NoError(t, err)
	err = createTestGoMod(parentDirectory)

	defer func() {
		err := os.RemoveAll(parentDirectory)
		require.NoError(t, err)
	}()

	moduleName, err := directoryModuleName(parentDirectory)
	require.NoError(t, err)
	require.Equal(t, testModuleName, moduleName)

	moduleName, err = directoryModuleName(childDirectory)
	require.NoError(t, err)
	require.Equal(t, testModuleName+"/child", moduleName)
}

func TestGoModModuleName(t *testing.T) {
	directoryName, err := ioutil.TempDir("", tempDirPattern)
	require.NoError(t, err)

	defer func() {
		err := os.RemoveAll(directoryName)
		require.NoError(t, err)
	}()

	err = createTestGoMod(directoryName)
	require.NoError(t, err)

	moduleName, err := goModModuleName(filepath.Join(directoryName, "go.mod"))
	require.Equal(t, moduleName, testModuleName)
}
