package main

import (
	"os"

	"git.adyanth.site/adyanth/shortpaste"
)

func main() {
	bind, ok := os.LookupEnv("SP_BIND_ADDR")
	if !ok {
		bind = ":8080"
	}

	storagePath, ok := os.LookupEnv("SP_STORAGE_PATH")
	if !ok {
		storagePath = "~/.shortpaste"
	}

	_, link307Redirect := os.LookupEnv("SP_307_REDIRECT")

	app := shortpaste.NewApp(bind, storagePath, link307Redirect)
	app.Run()
}
