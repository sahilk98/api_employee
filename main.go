package main

import (
	"fmt"
	"api_employee/database"
	api "api_employee/api/server"

)

func main() {

	errMigrate := database.MigrateTable()
	if errMigrate != nil {
		fmt.Println("Error while creating table: ", errMigrate)
	}
	fmt.Println("Table created")

	api.SetupRouterAPI()

}