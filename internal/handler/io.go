package handler

import (
	"errors"
	"github.com/dschemp/go-prntserve/internal/cmd"
	"io/ioutil"
	"os"
	"path"
	"syscall"
)

var (
	ErrNotADir        = errors.New("path is not a directory")
	ErrFileNotFound   = errors.New("file could not be found")
	folderPermissions = 0750
	probeFileName     = "__probe_"
)

func ProbeStoragePathOnFS() error {
	dir := cmd.StoragePath()

	// Check if directory exists
	stat, err := os.Stat(dir)
	// The error given by os.Stat should always be *PathError
	if pErr, ok := err.(*os.PathError); ok {
		if pErr.Err == syscall.ERROR_FILE_NOT_FOUND {
			// If the directory does not exist, attempt to create it
			if err := os.MkdirAll(dir, os.FileMode(folderPermissions)); err != nil {
				return err
			}
			// The directory we created could be created and thus does not need to be further probed
			return nil
		} else {
			// Return on any other error
			return pErr
		}
	} else if err != nil {
		// Panic if error is not *PathError as this should never happen
		panic(err)
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

func GetFileFromFS(fileName string) ([]byte, error) {
	fullPath := path.Join(cmd.FullStoragePath(), fileName)

	if !fileExistsOnFS(fullPath) {
		return nil, ErrFileNotFound
	}

	data, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func fileExistsOnFS(filePath string) bool {
	_, err := os.Stat(filePath)
	if pErr, ok := err.(*os.PathError); ok {
		return pErr.Err != syscall.ERROR_FILE_NOT_FOUND
	}

	return true
}
