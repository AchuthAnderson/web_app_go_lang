package main

import "gorm.io/gorm"


type Author struct {
	FullName string `json:"fullname"`
	Website string `json:"website"`
}

type Course struct {
	CourseId	string `json:"courseid"` 
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}


// User is a GORM model
type User struct {
	ID uint `gorm:"primaryKey"`
	UserName string `gorm:"unique"`
	Email string
}

type NewUser struct {
	gorm.Model
	FirstName string
	LastName string
	Email string
}