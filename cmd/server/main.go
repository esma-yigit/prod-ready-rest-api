package main

import "fmt"

//App => the struct which contains things like pointers
//to database connection
type App struct {
}

func (a *App) Run() error {
	fmt.Println("Setting up our App")
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
