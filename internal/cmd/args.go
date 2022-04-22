package cmd

import "flag"

var (
	port               int
	token              string
	showVersion        bool
	behindReverseProxy bool
)

func ParseArguments() {
	flag.IntVar(&port, "port", 8080, "Port number on which to listen")
	flag.StringVar(&token, "token", "CHANGE_ME", "Instance-wide bearer token for authorization")
	flag.BoolVar(&showVersion, "version", false, "Show version information")
	flag.BoolVar(&behindReverseProxy, "rproxy", false, "Is instance behind a reverse proxy which passes IP information?")
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
