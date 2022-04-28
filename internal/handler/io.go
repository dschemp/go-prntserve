package handler

import (
	"errors"
	"github.com/dschemp/go-prntserve/internal/cmd"
	"io/ioutil"
	"os"
	"path"
)

var (
	ErrNotADir           = errors.New("path is not a directory")
	ErrFileNotFound      = errors.New("file could not be found")
	ErrFileAlreadyExists = errors.New("file already exists")
	folderPermissions    = 0750
	filePermissions      = 0660
	probeFileName        = "__probe_"
)

// ProbeStoragePathOnFS tries to probe the target directory in which all files will be stored.
// This ensures that the directory is available at start of the app.
func ProbeStoragePathOnFS() error {
	dir := cmd.StoragePath()

	// Check if directory exists
	stat, err := os.Stat(dir)
	// The error given by os.Stat should always be *PathError
	if errors.Is(err, os.ErrNotExist) {
		// If the directory does not exist, attempt to create it
		if err := os.MkdirAll(dir, os.FileMode(folderPermissions)); err != nil {
			return err
		}
		// The directory we created could be created and thus does not need to be further probed
		return nil
	} else if err != nil {
		return err
	}

	if !stat.IsDir() {
		return ErrNotADir
	}

	// Try to create a temporary file
	tempFile, err := ioutil.TempFile(dir, probeFileName)
	if err != nil {
		return err
	}

	err = tempFile.Close()
	if err != nil {
		return err
	}

	err = os.Remove(tempFile.Name())
	if err != nil {
		return err
	}

	return nil
}

// GetFileFromStorage tries to retrieve a file by its relative path from the previously specified storage path.
// This can return ErrFileNotFound or any error returned by ioutil.ReadFile.
func GetFileFromStorage(relativeFilePath string) ([]byte, error) {
	fullPath := getAbsolutePathInStorage(relativeFilePath)

	if !FileExistsOnFS(fullPath) {
		return nil, ErrFileNotFound
	}

	data, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// SaveFileToStorage tries to save a file by its relative path to the previously specified storage path.
// It can return ErrFileAlreadyExists or any error returned by os.WriteFile.
func SaveFileToStorage(relativeFilePath string, data []byte) error {
	fullPath := getAbsolutePathInStorage(relativeFilePath)

	if FileExistsOnFS(fullPath) {
		return ErrFileAlreadyExists
	}

	err := os.WriteFile(fullPath, data, os.FileMode(filePermissions))
	if err != nil {
		return err
	}

	return nil
}

// FileExistsOnFS checks if a file exists on the file system.
// If the specified file exists, then true is returned. Otherwise, false.
func FileExistsOnFS(filePath string) bool {
	_, err := os.Stat(filePath)

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func getAbsolutePathInStorage(relativeFilePath string) string {
	return path.Join(cmd.FullStoragePath(), relativeFilePath)
}
