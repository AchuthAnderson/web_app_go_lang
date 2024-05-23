package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"encoding/json"
	"github.com/gorilla/mux"
)

const (
	contentType = "content-type"
	applicationJsonType = "application/json"
)

func postStatup() {
	bytes, err := json.Marshal(getAllCourses())
	if(err != nil) {
		panic("Failed to unmarshal getAllCourses "+ err.Error())
	}
	//WriteToMyDbFile(&bytes)
	createAndWriteToMyDb(&bytes)
	fmt.Println("Successfully wrote all courses to MyDb")
}

func main() {
	postStatup()
	fmt.Println("Welcome to Web App written in GO Lang")
	
	//Defining the Router
	r := mux.NewRouter()

	server := http.Server {
		Addr: ":8080",
		Handler: r,
	}

	//Routers
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", handlerGetAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", handlerGetCourseByID).Methods("GET")
	
	//Starting Server
	//log.Fatal(http.ListenAndServe(":4000", r))

	log.Fatal(server.ListenAndServe())
	time.Sleep(time.Second * 5)
	server.Shutdown(context.TODO())
}

