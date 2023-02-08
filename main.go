package main

import (
	"hua-proxy/client"
	"hua-proxy/server"
	"os"
	"strings"
)

func main() {
	env, b := os.LookupEnv("type")
	if b {
		if strings.EqualFold(env, "server") || strings.EqualFold(env, "SERVER") {
			server.Main()
		}
		if strings.EqualFold(env, "client") || strings.EqualFold(env, "CLIENT") {
			client.Main()
		}
	} else {
		server.Main()
	}
}
