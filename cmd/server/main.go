package main

import "fmt"

// App -contains things like pointers
// to DB connections
type App struct {
}

// Run - sets up our app
func (app *App) Run() error {
	fmt.Println("Setting up our app")
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
