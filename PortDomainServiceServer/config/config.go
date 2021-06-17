package config

import "os"

var VERSION = "v0.0.1"
var GitCommit = "Development"

var configItems = map[string]string{
	"SERVERPORT":  ":10000",
	"DatabaseDSN": "sqluser:sqluser@tcp(127.0.0.1:3306)/ports",
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
