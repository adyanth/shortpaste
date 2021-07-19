package main

import (
	"os"

	"git.adyanth.site/adyanth/shortpaste"
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

	_, link307Redirect := os.LookupEnv("LINK_307_REDIRECT")

	app := shortpaste.NewApp(bind, storagePath, link307Redirect)
	app.Run()
}
