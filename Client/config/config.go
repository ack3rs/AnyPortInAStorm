package config

import "os"

var VERSION = "v0.0.1"
var GitCommit = "Development"

var configItems = map[string]string{
	"RPC": ":10000",
}

// Get the Key.  First check the OS Env variables,  then the defaults used in development.
func Get(key string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}

	if configItems[key] != "" {
		return configItems[key]
	}
	return ""
}
