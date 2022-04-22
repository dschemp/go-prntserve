package handler

import (
    "errors"
    "github.com/dschemp/go-prntserve/internal/cmd"
    "io/ioutil"
    "os"
    "syscall"
)

var (
    ErrNotADir        = errors.New("path is not a directory")
    folderPermissions = 0750
    probeFileName     = "__probe_"
)

func ProbeStoragePath() error {
    path := cmd.StoragePath()

    // Check if directory exists
    stat, err := os.Stat(path)
    // The error given by os.Stat should always be *PathError
    if pErr, ok := err.(*os.PathError); ok {
        if pErr.Err == syscall.ERROR_FILE_NOT_FOUND {
            // If the directory does not exist, attempt to create it
            if err := os.MkdirAll(path, os.FileMode(folderPermissions)); err != nil {
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
    tempFile, err := ioutil.TempFile(path, probeFileName)
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
