package main

import (
	"fmt"
	"main/config"
	"main/database"
	"main/routes"
)

func main() {
	PORT := config.Config("DB_PORT")
	fmt.Printf("Server is Running at port %v \n", PORT)

	database.InitDatabase()

	routes.Run(PORT)

}
