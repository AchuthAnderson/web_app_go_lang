package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "log"
)

func connectToPostgresSql() (*gorm.DB, error) {
	dsn := "user=postgres password =password dbname=achuth-db host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

//CRUD operations using GORM
func createUser(db *gorm.DB, user *User) error {
	result := db.Create(user)
	if result.Error != nil {
		return  result.Error
	}
	return  nil
}

func getUserByID(db *gorm.DB, userId uint) (*User, error) {
	var user User
	result := db.First(&user, userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func updateUser(db *gorm.DB, user *User) error  {
	result := db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func deleteUser(db *gorm.DB, user *User) error {
	result := db.Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}