package server

import (
	"os"
	"syscall"
)

// Signal channel for initiating the shutdown.
var Signal = make(chan os.Signal)

// Fail reports an error and shuts down the server.
func Fail(err string, params ...interface{}) {
	if err != "" {
		log.Errorf(err, params...)
	}

	Shutdown()
}

// Shutdown gracefully stops the server.
func Shutdown() {
	Signal <- syscall.SIGINT
}
