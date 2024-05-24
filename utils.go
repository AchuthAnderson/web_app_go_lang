package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const myDbPath = "./myDb.json"

func ReadCoursesFromDbFile() []Course {
	contents, err := os.ReadFile(myDbPath)
	if err != nil {
		panic("Failed to Read file" + myDbPath + "with Error: "+err.Error())
	}
	var courses []Course
	json.Unmarshal(contents, &courses)
	return courses
}

func createAndWriteToMyDb(bytes *[]byte) {
	fmt.Println("Writing to myDb.json")
	fmt.Printf("Bytes received to write to file %s and no of bytes are: %d \n", string(*bytes), len(*bytes))
	err := os.WriteFile(myDbPath, *bytes, 0666)
	if err != nil {
		panic("Failed to write to file: "+ err.Error())
	}
	fmt.Println("Data successfully written to Json file")
}

//TODO: Write a function to append a course to myDB.json file
func addCourseToMyDb(newCourse Course) error {
	current_courses := ReadCoursesFromDbFile()
	current_courses = append(current_courses, newCourse)
	bytes, err := json.Marshal(current_courses)
	if err != nil {
		return errors.New("failed to add course to MyDB")
	}
	createAndWriteToMyDb(&bytes)
	return nil
}

//TODO: Write a fucntion to append multiple courses to myDB.json file. 

