package main

import (
	"fmt"
	"main/assets"
	"main/game"
	"os"
	"time"
)

func main() {

	resourcesPath := "resources"
	if len(os.Args) > 1 {
		resourcesPath = os.Args[1]
	}

	assets.GlobalAssetLibrary = assets.NewAssetLibrary()

	err := assets.GlobalAssetLibrary.ProcessFiles(resourcesPath)
	if err != nil {
		fmt.Printf("Error processing XML files: %s | %s", resourcesPath, err)
		return
	}

	addr := ":2050"
	server, err := game.NewAndServe(addr)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	game.GlobalServerListener = server
	defer game.GlobalServerListener.Stop()

	fmt.Println("Server started on", addr)

	game.GlobalWorldManager = game.NewWorldManager()
	game.GlobalWorldManager.CreateWorld("Nexus")

	running := true

	lastTime := time.Now()

	elapsed := 0.0

	for running {
		elapsedSeconds := time.Since(lastTime).Seconds()
		lastTime = time.Now()

		// every server frame
		server.ProcessConnectionMessages()

		// tick the server logic once every 200 ms
		elapsed += elapsedSeconds
		if elapsed >= 0.2 {
			game.GlobalWorldManager.TickAllWorlds(elapsed)
			elapsed = 0.0
		}
	}

	// handle cleanups

	fmt.Println("Server stopped")
}
