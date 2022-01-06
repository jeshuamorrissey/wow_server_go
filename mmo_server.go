package main

import (
	"sync"

	"github.com/jeshuamorrissey/wow_server_go/authserver"
	"github.com/jeshuamorrissey/wow_server_go/worldserver"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/config"
	"github.com/sirupsen/logrus"
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
		authserver.RunAuthServer(5000, config)
	}()

	go func() {
		defer wg.Done()
		worldserver.RunWorldServer("Sydney", 5001, config)
	}()

	wg.Wait()
}
