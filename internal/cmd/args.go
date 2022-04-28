package cmd

import (
	"flag"
	"github.com/go-chi/jwtauth"
	"os"
	"path"
)

var (
	port               int
	jwtSecret          string
	showVersion        bool
	behindReverseProxy bool
	storagePath        string
	debugLogging       bool
	useJSONLogging     bool
)

func ParseArguments() {
	flag.IntVar(&port, "port", 8080, "Port number on which to listen")
	flag.StringVar(&jwtSecret, "jwtSecret", "CHANGE_ME", "Instance-wide secret used for JWT authorization")
	flag.BoolVar(&showVersion, "version", false, "Show version information")
	flag.BoolVar(&behindReverseProxy, "rproxy", false, "Is instance behind a reverse proxy which passes IP information?")
	flag.StringVar(&storagePath, "storagePath", "files", "Folder path to the directory in which all files will be stored")
	flag.BoolVar(&debugLogging, "v", false, "Enable debug / verbose logging")
	flag.BoolVar(&useJSONLogging, "json", false, "Use JSON logging instead of logfmt")
	flag.Parse()
}

func Port() int {
	return port
}

func JWTSecret() string {
	return jwtSecret
}

func ShowVersion() bool {
	return showVersion
}

func BehindReverseProxy() bool {
	return behindReverseProxy
}

func StoragePath() string {
	return storagePath
}

func FullStoragePath() string {
	workDir, _ := os.Getwd()
	storageDir := path.Join(workDir, storagePath)
	return storageDir
}

func DebugLogging() bool {
	return debugLogging
}

func UseJSONLogging() bool {
	return useJSONLogging
}

func JWTTokenAuth() *jwtauth.JWTAuth {
	return jwtauth.New("HS256", []byte(JWTSecret()), nil)
}
