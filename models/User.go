package models

import (
	"gorm.io/gorm"
)

type User struct{
	gorm.Model

	FirstName 	string 	`json:"first_name" gorm:"not null"`
	SecondName 	string 	`json:"second_name"`
	LastNames 	string 	`json:"last_names" gorm:"not null"`
	Age 		int 	`json:"age"`
	Password  	string  `json:"password" gorm:"not null"`

}