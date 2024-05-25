package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

func getAllCourses() []Course {
	
	// courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 499,
	// 	Author: &Author{FullName: "Jay", Website: "loc.in"}})
	// courses = append(courses, Course{CourseId: "1", CourseName: "TypeJS", CoursePrice: 699,
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
	courses := getAllCourses()
	
	sort.Slice(courses , func(i int, j int) bool {
		first, _ := strconv.Atoi(courses[i].CourseId)
		second, _ := strconv.Atoi(courses[j].CourseId)
		return first  < second
	})

	json.NewEncoder(w).Encode(courses)
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

func handlerCreateCourse(w http.ResponseWriter,r *http.Request) {
	
	w.Header().Set(contentType, applicationJsonType)
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Request Body is missing")
	}

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		panic("Failed to Decode RequestBody with Error: " + err.Error())
	}
	err = addCourseToMyDb(course)
	if err!=nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Failed to add Course")
	}
	w.WriteHeader(http.StatusCreated)
}

func handlerDeleteCourse(w http.ResponseWriter, r* http.Request) {
	w.Header().Set(contentType, applicationJsonType)
	params :=mux.Vars(r)
	courses := getAllCourses()
	courseIdToDelete := params["id"]
	
	fmt.Printf("Course with ID: %s will be deleted \n", courseIdToDelete)
	
	for i, course := range courses {
		if course.CourseId == courseIdToDelete {
			courses = append(courses[:i], courses[i+1:]...)
			bytes, _ := json.Marshal(courses)
			createAndWriteToMyDb(&bytes)
			w.WriteHeader(http.StatusNoContent)
			break
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("Id requested to delete doesn't exists")
}

func handleUpdateCourse(w http.ResponseWriter, r* http.Request) {
	
	w.Header().Set(contentType, applicationJsonType)
	params := mux.Vars(r)
	courseIdTobeUpdated := params["id"]
	
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Request Body is missing")
	}

	courses := getAllCourses()
	for index, course := range courses {
		if course.CourseId == courseIdTobeUpdated {
			courses = append(courses[:index], courses[index+1:]...)
			var newCourse Course
			json.NewDecoder(r.Body).Decode(&newCourse)
			newCourse.CourseId = courseIdTobeUpdated
			courses = append(courses, newCourse)

			bytes, err := json.Marshal(courses)
			if err != nil {
				panic("Failed to marshell courses with error: " +err.Error())
			}
			createAndWriteToMyDb(&bytes)
			w.WriteHeader(http.StatusOK)
			break
		}
	} 

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("Failed to update course")
}