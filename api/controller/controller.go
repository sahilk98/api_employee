package api

import (
	model "api_employee/api/model"
	module "api_employee/api/module"
	//"fmt"
	"strconv"

	//"database/sql"
	database "api_employee/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Endpoint 1 : Get employees
func GetEmployeeData(c *gin.Context){

	dbConnection, errConnect := database.ConnectDB()
	if errConnect != nil {
		HandleErrorEmployeeData(c , 500, errConnect.Error())
		return
	}

	defer dbConnection.Close()

	data, errGetData := module.GetEmployeeDataModule(dbConnection)
	if errGetData != nil {
		HandleErrorEmployeeData(c, 400, errGetData.Error())
		return
	}

	HandleSuccessGetEmployeeData(c, data)
}

func HandleErrorEmployeeData(c *gin.Context, errorCode int, message string){
	var errorResp model.RespErrorEmployeeData
	errorResp.Code = errorCode
	errorResp.Message = message

	c.JSON(errorCode, errorResp)
}

func HandleSuccessGetEmployeeData(c *gin.Context, data []model.RespEmployeeAllData){
	c.JSON(http.StatusOK, data)
}

// Endpoint 2 : Get employee Data by ID

func GetEmployeeDataByID(c *gin.Context){
	dbConnection, errConnect := database.ConnectDB()
	if errConnect != nil {
		HandleErrorEmployeeData(c , 500, errConnect.Error())
		return
	}

	defer dbConnection.Close()

	idStr := c.Param("id")
	id, errConvString := strconv.Atoi(idStr)
	if errConvString != nil{
		HandleErrorEmployeeData(c, 400, "Error converting ID value to integer: " + errConvString.Error())
		return
	}


	data, errGetData := module.GetEmployeeDataByID(id,dbConnection)
	if errGetData != nil {
		HandleErrorEmployeeData(c, 400, errGetData.Error())
		return
	}

	HandleSuccessGetEmployeeDataByID(c, data)
}

func HandleSuccessGetEmployeeDataByID(c *gin.Context, data model.RespEmployeeAllData){
	c.JSON(http.StatusOK, data)
}



// Endpoint 3 : POST New Employee

func PostNewEmployee(c *gin.Context){
	dbConnection, errConnect := database.ConnectDB()
	if errConnect != nil {
		HandleErrorEmployeeData(c , 500, errConnect.Error())
		return
	}

	defer dbConnection.Close()

	requestData := model.ReqPostNewEmployeeData{}

	errBindJson := c.BindJSON(&requestData)
	if errBindJson != nil{
		HandleErrorEmployeeData(c, 400, errBindJson.Error())
		return
	}

	message, errInsertData := module.PostNewEmployeeModule(requestData, dbConnection)
	if errInsertData != nil{
		HandleErrorEmployeeData(c, 400, errInsertData.Error())
		return
	}

	HandleSuccessPostEmployeeData(c, message)

}

func HandleSuccessPostEmployeeData(c *gin.Context, message string){
	var response model.RespSuccessPostData 
	response.Message = message 

	c.JSON(http.StatusOK, response)
}

// Endpoint 4 : Update Employee record using ID by PUT

func UpdateEmployeeDataByID(c *gin.Context){
	dbConnection, errConnect := database.ConnectDB()
	if errConnect != nil {
		HandleErrorEmployeeData(c , 500, errConnect.Error())
		return 
	}

	defer dbConnection.Close()

	idStr := c.Param("id")
	id, errConvString := strconv.Atoi(idStr)
	if errConvString != nil{
		HandleErrorEmployeeData(c, 400, "Error converting ID value to integer: " + errConvString.Error())
		return 
	}

	if idStr == ""{
		HandleErrorEmployeeData(c,400, "ID cannot be empty")
		return 
	}

	requestData := model.ReqUpdateEmployeeByID{}
	errBindJson := c.BindJSON(&requestData)
	if errBindJson != nil{
		HandleErrorEmployeeData(c, 400, errBindJson.Error())
	}

	message, errUpdateData := module.UpdateEmployeeDataByID(requestData, id, dbConnection)
	if errUpdateData != nil{
		HandleErrorEmployeeData(c,400, "Error Updating Employee Data :"+ errUpdateData.Error())
		return 
	}

	HandleSuccessUpdateEmployeeData(c,message)

}

func HandleSuccessUpdateEmployeeData(c *gin.Context, message string){
	var response model.RespSuccessPostData 
	response.Message = message 

	c.JSON(http.StatusOK, response)
}

// Endpoint 5 : Delete employee data by ID
func DeleteEmployeeDataByID(c *gin.Context){
	dbConnection, errConnect := database.ConnectDB()
	if errConnect != nil {
		HandleErrorEmployeeData(c , 500, errConnect.Error())
		return 
	}

	defer dbConnection.Close()

	idStr := c.Param("id")
	id, errConvString := strconv.Atoi(idStr)
	if errConvString != nil{
		HandleErrorEmployeeData(c, 400, "Error converting ID value to integer: " + errConvString.Error())
		return 
	}
	if idStr == ""{
		HandleErrorEmployeeData(c,400, "ID cannot be empty")
		return 
	}
	message, errDeleteData := module.DeleteEmployeeDataByID(id, dbConnection)
	if errDeleteData != nil{
		HandleErrorEmployeeData(c,400, "Error Deleting Employee Data :"+ errDeleteData.Error())
		return 
	}
	HandleSuccessDeleteEmployeeData(c,message)

}

func HandleSuccessDeleteEmployeeData(c *gin.Context, message string){
	var response model.RespSuccessPostData 
	response.Message = message 

	c.JSON(http.StatusOK, response)
}
