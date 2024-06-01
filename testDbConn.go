package main

import (
	"fmt"
	"log"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
)

func testGORMDBConnection() bool {
	fmt.Println("Testing DB connecting using GROM on User table.")

	db, err := connectToPostgresSql()
	if err!= nil {
		log.Fatal("Error occured while connecting to DB")
		return false
	}

	user := User{UserName : "AchuthAnderson", Email:"AndersonReddy@duck.com"}
	err = createUser(db, &user)
	if err != nil {
		log.Fatal("Error occurred while adding User to user table: ", err.Error())
		return false
	}
	fmt.Println("User Successfully added to user table")
	
	queriedUser , err := getUserByID(db, user.ID)
	if err != nil {
		log.Fatal("Error occurred while querying User to user table: ", err.Error())
		return false
	}

	if user != *queriedUser {
		log.Fatal("Wrong user fetched from DB")
		return false
	}
	fmt.Println("User successfully fetched from DB")
	user.UserName = "Anderson_Reddy"
	err = updateUser(db, &user)
	if err != nil {
		log.Fatal("Error occurred while updating User: ", err.Error())
		return false
	}
	fmt.Println("User updated successfully")
	err = deleteUser(db, &user)
	if err!= nil {
		log.Fatal("Failed to delete user from DB: ", err.Error())
		return false
	}
	fmt.Println("user deleted from DB successfully")
	_ , err = getUserByID(db, user.ID)
	if err == nil {
		log.Fatal("above delete query failed to delete user from user table")
		return false
	}
	fmt.Println("Test run successfully")
	return true
}

