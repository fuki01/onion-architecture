package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	User   string
	Pass   string
	Host   string
	DBName string
}

func NewDatabase(user, pass, host, dbname string) *Database {
	return &Database{
		User:   user,
		Pass:   pass,
		Host:   host,
		DBName: dbname,
	}
}

func (d *Database) Connect() (*gorm.DB, error) {
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", d.User, d.Pass, d.Host, d.DBName)
	db, err := connectToDatabase(connection)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDatabase(connection string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	maxAttempts := 5
	interval := time.Second * 2

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		db, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
		if err == nil {
			return db, nil
		}
		log.Printf("Failed to connect to database (attempt %d): %v", attempts, err)
		time.Sleep(interval)
	}

	return nil, fmt.Errorf("failed to connect to database after %d attempts", maxAttempts)
}
