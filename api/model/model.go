package api

import "time"

type (
	RespEmployeeAllData struct {
		ID        int       `json:"id"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Email     string    `json:"email"`
		HireDate  time.Time `json:"hire_date"`
	}

	RespErrorEmployeeData struct{
		Code int `json:"code"`
		Message string `json:"message"`
	}

	ReqPostNewEmployeeData struct{
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Email     string    `json:"email"`
		HireDate  string    `json:"hire_date"`
	}

	RespSuccessPostData struct{
		Message string `json:"message"`
	}

	ReqUpdateEmployeeByID struct{
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Email     string    `json:"email"`
	}
)