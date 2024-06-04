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
		log.Fatal("Failed to unmarshal getAllCourses ", err.Error())
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
	r.HandleFunc("/courses", handlerCreateCourse).Methods("POST")	
	r.HandleFunc("/courses/id", handleUpdateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", handlerDeleteCourse).Methods("Delete")
	
	//Starting Server
	//log.Fatal(http.ListenAndServe(":4000", r))
	//connectToPostgresUsingGORM()
	//testGORMDBConnection()

	//db, _ := connectToPostgresSql()
	//db.AutoMigrate(&NewUser{})
	//conn := GetDbConn()
	//GetAllAuthors(conn)
	//GetAllCourses(conn)
	//AlterTableCourseToHaveUniqueCourseId(conn)
	//demoTransaction(conn)

	createJWT("admin")
	log.Fatal(server.ListenAndServe())
	time.Sleep(time.Second * 5)
	server.Shutdown(context.TODO())
}

func connectToPostgresUsingGORM() {
	dbConn , err := connectToPostgresSql()
	if err!=nil {
		log.Fatal("Error occured while connecting to postgreSql from GORM: ", err.Error())
	}

	db, _ := dbConn.DB()
	err = db.Ping()
	if err!=nil {
		log.Fatal("Failed to ping Db with error: ", err.Error())
	}
	fmt.Println("Ping successfull..")

	err = dbConn.AutoMigrate(&User{})
	if err!=nil {
		log.Fatal("AutoMigration failed with Error: ", err.Error())
	}
	fmt.Println("Auto Migration successful")
}