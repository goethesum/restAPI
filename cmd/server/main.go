package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/goethesum/restAPI/internal/transport/http"
)

// App -contains things like pointers
// to DB connections
type App struct {
}

// Run - sets up our app
func (app *App) Run() error {
	fmt.Println("Setting up our app")

	handler := transportHTTP.NewHandler()
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
