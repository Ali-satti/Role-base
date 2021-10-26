package models

type UserRole struct {
	Designation string `json:"designation" structs:"designation"`
	ApiName string `json:"api_name" structs:"api_name"`
}
