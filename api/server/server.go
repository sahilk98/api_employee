package api

import (
	controllerEmployeeData "api_employee/api/controller"
	// "database/sql"
	// "fmt"

	"github.com/gin-gonic/gin"
)

type AppServer struct{
	port string 
	engine *gin.Engine
}

// func NewAppServer(portVal string) *AppServer{
// 	return &AppServer{
// 		port: portVal,
// 		engine: gin.New(),
// 	}
// }

func SetupRouterAPI(){


	router := gin.Default()

	router.GET("/employees", controllerEmployeeData.GetEmployeeData)
	router.GET("/employees/:id", controllerEmployeeData.GetEmployeeDataByID)
	router.POST("/employees", controllerEmployeeData.PostNewEmployee)
	router.PUT("/employees/:id", controllerEmployeeData.UpdateEmployeeDataByID)
	router.DELETE("/employees/:id", controllerEmployeeData.DeleteEmployeeDataByID)


	router.Run(":8080")

}


