package main

import (
	"fmt"
	"net/http"

	"github.com/goethesum/restAPI/internal/comment"
	"github.com/goethesum/restAPI/internal/database"
	transportHTTP "github.com/goethesum/restAPI/internal/transport/http"
)

// App -contains things like pointers
// to DB connections
type App struct {
}

// Run - sets up our app
func (app *App) Run() error {
	fmt.Println("Setting up our app")

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
	fmt.Println("Failed to set up server")
	return nil
}

func main() {
	fmt.Println("behold")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up REST API")
		fmt.Println(err)
	}
}
