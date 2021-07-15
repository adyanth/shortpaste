package shortpaste

import (
	"fmt"
	"os/user"
	"path"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (app *App) Run() {
	app.db.AutoMigrate(&Link{})
	app.db.AutoMigrate(&File{})
	app.db.AutoMigrate(&Text{})
	fmt.Println("Migration complete")
	app.handleRequests()
}

func NewApp(bind, dbUri, storagePath string) App {
	usr, _ := user.Current()
	if storagePath == "~" {
		storagePath = usr.HomeDir
	} else if strings.HasPrefix(storagePath, "~/") {
		storagePath = path.Join(usr.HomeDir, storagePath[2:])
	}

	if db, err := gorm.Open(sqlite.Open(dbUri), &gorm.Config{}); err == nil {
		return App{
			bind:        bind,
			db:          db,
			storagePath: storagePath,
		}
	} else {
		panic(fmt.Errorf("db error %v", err))
	}
}
