package main

import (
	"fmt"
	"github.com/esma-yigit/prod-ready-rest-api/internal/database"
	transportHTTP "github.com/esma-yigit/prod-ready-rest-api/internal/transport/http"
	"net/http"
)

//App => the struct which contains things like pointers
//to database connection
type App struct {
}

func (a *App) Run() error {
	fmt.Println("Setting up our App")
	var err error
	_, err = database.NewDatabase()
	if err != nil {
		return err
	}

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

//Run => sets up our application

func main() {
	fmt.Println("Go REST API Course")
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our application")
		fmt.Println(err)
	}
}
