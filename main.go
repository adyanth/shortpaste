package main

import (
	"git.adyanth.site/adyanth/shortpaste/shortpaste"
)

func main() {
	app := shortpaste.NewApp(":8080", "./test.db", "./")
	app.Run()
}
