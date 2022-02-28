package main

import (
	"fmt"
	"net/http"

	"github.com/goethesum/restAPI/internal/comment"
	"github.com/goethesum/restAPI/internal/database"
	transportHTTP "github.com/goethesum/restAPI/internal/transport/http"

	log "github.com/siruspen/logrus"
)

// App - contain aplication information
type App struct {
	Name    string
	Version string
}

// Run - sets up our app
func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting up application")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}
	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		return err
	}
	log.Error("Failed to set up server")
	return nil
}

func main() {
	fmt.Println("behold")
	app := App{
		Name:    "Rest-API",
		Version: "1.0.0",
	}
	if err := app.Run(); err != nil {
		log.Error("Error starting up REST API")
		log.Fatal(err)
	}
}
