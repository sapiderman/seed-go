package main

import (
	"fmt"

	"github.com/sapiderman/seed-go/internal"
	log "github.com/sirupsen/logrus"
)

var (
	spashScreen = `
	🇸​​​​​🇪​​​​​🇷​​​​​🇻​​​​​🇪​​​​​🇷​​​​​ 🇺​​​​​🇵​​​​​
	
	golang seed-go
	https://github.com/sapiderman/seed-go/blob/master/README.md
	`
)

func init() {
	fmt.Println(spashScreen)
	log.Info("intializing")
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

// Main entry point
func main() {

	// start server
	internal.StartServer()
}
