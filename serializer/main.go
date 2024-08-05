package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	Name        []byte                 `gorm:"serializer:json"`
	Roles       Roles                  `gorm:"serializer:json"`
	Contracts   map[string]interface{} `gorm:"serializer:json"`
	JobInfo     Job                    `gorm:"type:bytes;serializer:gob"`
	CreatedTime int64                  `gorm:"serializer:unixtime;type:time"` // store int as datetime into database
	gorm.Model
}

type Roles []string

type Job struct {
	Title      string
	Location   string
	IsIntern   bool
	Experience int
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	slog.Info("server started .....")

	createdAt := time.Date(2024, 1, 1, 0, 8, 0, 0, time.UTC)
	data := User{
		Name:        []byte("Jenny Gutierez"),
		Roles:       []string{"admin", "owner"},
		Contracts:   map[string]interface{}{"name": "Jenny Gutierez", "age": 28, "experience": 4},
		CreatedTime: createdAt.Unix(),
		JobInfo: Job{
			Title:      "Developer",
			Location:   "London",
			IsIntern:   false,
			Experience: 4,
		},
	}
	db, err := dbconnect()
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info("server stared ...")
	err = db.Create(&data).Error
	if err != nil {
		slog.Warn("failed to create user object ")
	}
	var result User
	db.First(&result, "id = ?", data.ID)
	fmt.Println(strings.Repeat("+", 50))
	fmt.Printf("user %#v \n", result)
	fmt.Println(strings.Repeat("+", 50))
}
func dbconnect() (*gorm.DB, error) {
	slog.Info("connecting to db ...")
	db, err := gorm.Open(sqlite.Open("serializer.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate %w", err)
	}
	return db, nil
}
