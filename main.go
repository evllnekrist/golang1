package main

import (
	"photo-app/database"
	"photo-app/router"
	"os"
)

func main() {
	database.StartDB()
	var port = os.Getenv("PORT")
	router.StartApp().Run(":"+port)
	// r := router.StartApp()
	// r.Run(":8080")
}
