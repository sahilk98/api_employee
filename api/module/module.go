package api

import (
	model "api_employee/api/model"
	"database/sql"
	"strconv"

	//"encoding/json"
	"fmt"
)

// Endpoint 1 : Get Employee Data
func GetEmployeeDataModule(db *sql.DB)([]model.RespEmployeeAllData, error){
	query := "SELECT id, first_name, last_name, email, hire_date FROM employees; "

	rows, errSelect := db.Query(query)
	if errSelect != nil {
		return []model.RespEmployeeAllData{}, fmt.Errorf("error while executing query: ", errSelect)
	}
	defer rows.Close()

	var data []model.RespEmployeeAllData
	for rows.Next(){
		var dataEachRow model.RespEmployeeAllData
		errScan := rows.Scan(&dataEachRow.ID, &dataEachRow.FirstName, &dataEachRow.LastName, &dataEachRow.Email, &dataEachRow.HireDate)
		if errScan != nil {
			return []model.RespEmployeeAllData{}, fmt.Errorf("error while scanning result into json struct: ", errScan)
		}
		data = append(data, dataEachRow)
	}

	return data, nil

}



// Endpoint 2 : Get employee by ID using GET

func GetEmployeeDataByID(id int, db *sql.DB) (model.RespEmployeeAllData, error){

	queryToPrepare := "SELECT id, first_name, last_name, email, hire_date FROM employees WHERE id = $1 ;"
	selectQuery, errPrepareQuery := db.Prepare(queryToPrepare)
	if errPrepareQuery != nil {
		return model.RespEmployeeAllData{}, fmt.Errorf("error while preparing select query: ", errPrepareQuery.Error())
	}
	defer selectQuery.Close()

	var data model.RespEmployeeAllData

	errSelect := selectQuery.QueryRow(id).Scan(&data.ID, &data.FirstName, &data.LastName, &data.Email, &data.HireDate)
	if errSelect != nil {
		return model.RespEmployeeAllData{}, fmt.Errorf("error while executing select based on ID: ", errSelect.Error())
	}

	return data, nil

}



// Endpoint 3 : Insert New Employee Data
func PostNewEmployeeModule(newData model.ReqPostNewEmployeeData, db *sql.DB)(string, error){

	queryToPrepare := " INSERT INTO employees (first_name, last_name, email, hire_date) VALUES ($1, $2, $3, $4); "

	insertQuery, errPrepareQuery := db.Prepare(queryToPrepare)
	if errPrepareQuery != nil {
		return "", fmt.Errorf("error while preparing insert statement: ", errPrepareQuery.Error())
	}
	defer insertQuery.Close()

	_, errExec := insertQuery.Exec(newData.FirstName, newData.LastName, newData.Email, newData.HireDate)
	if errExec != nil {
		return "", fmt.Errorf("error while executing insert script: ", errExec.Error())
	}

	return "Data successfully inserted", nil

}

// Endpoint 4 : Update Employee Data By ID
func UpdateEmployeeDataByID(newData model.ReqUpdateEmployeeByID  ,id int, db *sql.DB)(string, error){
	queryToPrepare := " UPDATE employees SET first_name = $1, last_name = $2, email = $3 WHERE id = $4 ;"

	updateQuery, errPrepareQuery := db.Prepare(queryToPrepare)
	if errPrepareQuery != nil {
		return "", fmt.Errorf("error while preparing update statement: ", errPrepareQuery.Error())
	}
	defer updateQuery.Close()

	_, errExec := updateQuery.Exec(newData.FirstName, newData.LastName, newData.Email, id)
	if errExec != nil {
		return "", fmt.Errorf("error while executing update script: ", errExec.Error())
	}

	idStr := strconv.Itoa(id)

	return "Successfully updated data for ID : "+ idStr , nil

}

// Endpoint 5 : Delete employee Data by ID
func DeleteEmployeeDataByID(id int, db *sql.DB)(string, error){
	queryToPrepare := " DELETE FROM employees WHERE id = $1 ;"

	deleteQuery, errPrepareQuery := db.Prepare(queryToPrepare)
	if errPrepareQuery != nil {
		return "", fmt.Errorf("error while preparing update statement: ", errPrepareQuery.Error())
	}
	defer deleteQuery.Close()

	_, errExec := deleteQuery.Exec(id)
	if errExec != nil {
		return "", fmt.Errorf("error while executing update script: ", errExec.Error())
	}

	idStr := strconv.Itoa(id)

	return "Successfully deleted data for ID : "+ idStr , nil
}