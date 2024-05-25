package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)
/*
	- TODO:
		- Connecting to DB
			- Using Connection string
			- ConnectionPool
				- Default Properties 
				- How to modify connection pool properties. 
		- Rows
			- Map them to Struct
			- Check Name and type of the column 
			- Transaction
				- Create
				- Start
				- End
				- Rollover.
		- Create a table in DB 
		- Create a DB in the application. 
			- Read  current DB name from Application. 
*/

const DbConnStr = "postgres://postgres:password@localhost:5432/achuth-db?sslmode=disable"

func GetDbConn() *sql.DB {
	db, err := sql.Open("postgres", DbConnStr)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("Connect is successfull")
	return db
}

func ConnectToDB() {
	//connStr := "user:postgres dbname=achuth-db sslmode=verify-full"
	connStr := "postgres://postgres:password@localhost:5432/achuth-db?sslmode=disable"
	fmt.Println("Connection string is : ", connStr)
	db, err := sql.Open("postgres", connStr)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("Connect is successfull")
	err = db.Ping()
	if err != nil {
		fmt.Println("Error occured during ping: ", err.Error())
	}else{
		fmt.Println("Ping is successfull")
	}

	//var row sql.Rows
	rows , err := db.Query("select * from customers")
	if err!=nil {
		log.Fatal("Error occured while querying db: ", err.Error())
	}
	columnTypes, _ := rows.ColumnTypes()

	fmt.Println("Num of columns: ", len(columnTypes))
	for i, columnType := range columnTypes {
		fmt.Printf("Column Index is %d and ColumnType is: %s and %v", i, columnType.Name(), *columnType)
	}
}

func GetAllAuthors(conn *sql.DB) {
	fmt.Println("Reading all Authors from DB")
	rows, err := conn.Query("select full_name, website from author")
	if err != nil {
		log.Fatalln("Failed to read all records from authors table with err: ", err.Error())
	}

	var authors []Author
	for rows.Next() {
		var author Author
		err := rows.Scan(&author.FullName, &author.Website)
		if err != nil {
			log.Fatal("Failed to convert row into author : ", err.Error())
		}
		authors = append(authors, author)
	}
	rows.Close()
	fmt.Println("All authors from DB are :", authors)
}

func GetAllCourses(conn *sql.DB) {
	fmt.Println("Reading all courses from DB")
	var courses []Course
	rows, err := conn.Query("select course_id, course_name, course_price from course")
	
	if err != nil {
		log.Fatalln("Failed to read all records from course table with err: ", err.Error())
	}

	for rows.Next() {
		var course Course
		err := rows.Scan(&course.CourseId, &course.CourseName, &course.CoursePrice)
		if err != nil {
			log.Fatal("Failed to convert row into course : ", err.Error())
		}
		courses = append(courses, course)
	}
	rows.Close()
	fmt.Println("All courses from DB are :", courses)
}


