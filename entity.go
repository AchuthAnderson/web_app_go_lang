package main


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