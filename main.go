package main

import (
	"github.com/ssss.tantalum/my-restful-api/database"
	"github.com/ssss.tantalum/my-restful-api/router"
)

func main() {
	database.Connect("root:password@tcp(db)/my_database?parseTime=true")
	database.Migrate()

	router := router.InitRouter()
	router.Run(":8080")
}
