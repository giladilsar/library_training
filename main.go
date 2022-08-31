package main

import (
	"gin/config"
	"gin/server"
)

func main() {
	config.Setup()
	router := server.SetupRoutes()
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
