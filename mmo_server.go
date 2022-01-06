package main

import (
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/jeshuamorrissey/wow_server_go/lib/config"
	"github.com/jeshuamorrissey/wow_server_go/server/auth"
	"github.com/jeshuamorrissey/wow_server_go/server/world"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetLevel(logrus.DebugLevel)

	// Load the world config and object manager.
	// TODO(jeshua): make this load from a command-line flag maybe?
	config := config.NewConfigFromJSON("world.json")

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		auth.RunAuthServer(5000, config)
	}()

	go func() {
		defer wg.Done()
		world.RunWorldServer("Sydney", 5001, config)
	}()

	wg.Wait()
}
