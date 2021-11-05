package shortpaste

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Run migrates the DB and starts the web server
func (app *App) Run() {
	app.db.AutoMigrate(&Link{})
	app.db.AutoMigrate(&File{})
	app.db.AutoMigrate(&Text{})
	fmt.Println("Migration complete")
	app.handleRequests()
}

// NewApp creates a new App instance with the provided bind address and storage path
func NewApp(bind, storagePath, username, password string, noAuth, link307Redirect bool) App {
	usr, _ := user.Current()
	if storagePath == "~" {
		storagePath = usr.HomeDir
	} else if strings.HasPrefix(storagePath, "~/") {
		storagePath = path.Join(usr.HomeDir, storagePath[2:])
	}
	os.MkdirAll(path.Join(storagePath, "db"), 0700)

	dbUri := path.Join(storagePath, "db", "mapping.db")
	if db, err := gorm.Open(sqlite.Open(dbUri), &gorm.Config{}); err != nil {
		panic(fmt.Errorf("db error %v", err))
	} else {
		return App{
			bind:            bind,
			db:              db,
			storagePath:     storagePath,
			link307Redirect: link307Redirect,
			username:        username,
			password:        password,
			noAuth:          noAuth,
		}
	}
}
