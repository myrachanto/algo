package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
	Age  string
	Post []Post
}

type Post struct {
	ID       uint
	UserID   uint
	Content  string
	Comments []Comment
}

type Comment struct {
	ID     uint
	PostID uint
	Text   string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	// Insert sample data
	users := []User{
		{Name: "Alice", Age: "25"},
		{Name: "Bob", Age: "30"},
		{Name: "Charlie", Age: "35"},
	}

	for _, user := range users {
		db.Create(&user)
	}

	posts := []Post{
		{UserID: 1, Content: "Alice's first post"},
		{UserID: 1, Content: "Alice's second post"},
		{UserID: 2, Content: "Bob's post"},
		{UserID: 3, Content: "Charlie's post"},
	}

	for _, post := range posts {
		db.Create(&post)
	}

	comments := []Comment{
		{PostID: 1, Text: "First comment on Alice's first post"},
		{PostID: 2, Text: "First comment on Alice's second post"},
		{PostID: 3, Text: "First comment on Bob's post"},
		{PostID: 4, Text: "First comment on Charlie's post"},
	}

	for _, comment := range comments {
		db.Create(&comment)
	}

	// Call backtracking function
	var result []User
	findUsersByPostAndComment(db, 0, []User{}, &result, "First comment on Alice's first post")
	fmt.Println(result)
}

func findUsersByPostAndComment(db *gorm.DB, start int, path []User, result *[]User, commentText string) {
	var users []User
	db.Preload("Post.Comments").Find(&users)

	if start == len(users) {
		return
	}

	for i := start; i < len(users); i++ {
		for _, post := range users[i].Post {
			for _, comment := range post.Comments {
				if comment.Text == commentText {
					newPath := append(path, users[i])
					*result = append(*result, newPath...)
					findUsersByPostAndComment(db, i+1, newPath, result, commentText)
				}
			}
		}
	}
}
