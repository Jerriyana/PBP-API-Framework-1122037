package main

import (
	"fmt"
	"log"
	"martini/controllers"
	"net/http"

	"github.com/codegangsta/martini"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	m := martini.Classic()

	// 1. Routing untuk endpoint GET
	m.Get("/users", controllers.GetAllUsers)

	// 2. Routing untuk endpoint POST
	m.Post("/users", controllers.InsertUser)

	// 3. Routing untuk endpoint PUT
	m.Put("/users/:id", controllers.UpdateUser)

	// 4. Routing untuk endpoint DELETE
	m.Delete("/users/:id", controllers.DeleteUser)

	// Connected
	fmt.Println("Connected to port 8890")
	log.Println("Connected to port 8890")
	log.Fatal(http.ListenAndServe(":8890", m))
}
