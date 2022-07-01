package cmd

import (
	"github.com/go-chi/jwtauth"
	"os"
	"path"
)

type settings struct {
	ListenAddress        string
	JWTSecret            string
	IsBehindReverseProxy bool
	StoragePath          string
	UseDebugLogging      bool
	UseJSONLogging       bool
}

var SETTINGS settings

func ListenAddress() string {
	return SETTINGS.ListenAddress
}

func JWTSecret() string {
	return SETTINGS.JWTSecret
}

func IsBehindReverseProxy() bool {
	return SETTINGS.IsBehindReverseProxy
}

func StoragePath() string {
	return SETTINGS.StoragePath
}

func FullStoragePath() string {
	workDir, _ := os.Getwd()
	storageDir := path.Join(workDir, StoragePath())
	return storageDir
}

func DebugLogging() bool {
	return SETTINGS.UseDebugLogging
}

func UseJSONLogging() bool {
	return SETTINGS.UseJSONLogging
}

func JWTTokenAuth() *jwtauth.JWTAuth {
	return jwtauth.New("HS256", []byte(JWTSecret()), nil)
}
