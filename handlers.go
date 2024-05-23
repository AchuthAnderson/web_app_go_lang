package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getAllCourses() []Course {
	
	// courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 499,
	// 	Author: &Author{FullName: "Jay", Website: "loc.in"}})
	// courses = append(courses, Course{CourseId: "3", CourseName: "TypeJS", CoursePrice: 699,
	// 	Author: &Author{FullName: "Jk", Website: "loc.in"}})
	// return &courses
	
	return ReadCoursesFromDbFile()
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is my first API building"))
}

func handlerGetAllCourses(w http.ResponseWriter, r* http.Request) {

	bytes, err := json.Marshal(getAllCourses())
	
	if err!=nil {
		fmt.Println("Failed to marshall courses", err.Error())
	}else {
		coursesStr := string(bytes)
		fmt.Println("Courses as String: ", coursesStr)
	}
	
	// w.WriteHeader(http.StatusCreated)
	// w.Write(bytes)
	w.Header().Set(contentType,applicationJsonType)	
	json.NewEncoder(w).Encode(getAllCourses())
}

func handlerGetCourseByID(w http.ResponseWriter, r *http.Request) {
	 
	params := mux.Vars(r)

	for _, course := range getAllCourses() {
		if course.CourseId == params["id"] {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Requested Course Id is not found"))
}