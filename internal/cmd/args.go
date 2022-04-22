package cmd

import "flag"

var (
	port               int
	token              string
	showVersion        bool
	behindReverseProxy bool
	storagePath        string
)

func ParseArguments() {
	flag.IntVar(&port, "port", 8080, "Port number on which to listen")
	flag.StringVar(&token, "token", "CHANGE_ME", "Instance-wide bearer token for authorization")
	flag.BoolVar(&showVersion, "version", false, "Show version information")
	flag.BoolVar(&behindReverseProxy, "rproxy", false, "Is instance behind a reverse proxy which passes IP information?")
	flag.StringVar(&storagePath, "storagePath", "files", "Folder path to the directory in which all files will be stored")
	flag.Parse()
}

func Port() int {
	return port
}

func Token() string {
	return token
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
