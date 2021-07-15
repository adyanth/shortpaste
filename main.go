package main

import (
	"os"

	"git.adyanth.site/adyanth/shortpaste/shortpaste"
)

func main() {
	bind, ok := os.LookupEnv("BIND_ADDR")
	if !ok {
		bind = ":8080"
	}

	storagePath, ok := os.LookupEnv("STORAGE_PATH")
	if !ok {
		storagePath = "~/.shortpaste"
	}

	app := shortpaste.NewApp(bind, storagePath)
	app.Run()
}
