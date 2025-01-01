package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Define structures to hold JSON responses for posts, comments, and users.
type Post struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
}

type Comment struct {
	PostID    int    `json:"postId"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Body      string `json:"body"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func fetchPosts(wg *sync.WaitGroup) {
	defer wg.Done() // Decrements the counter when the goroutine completes
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Println("Error fetching posts:", err)
		return
	}
	defer resp.Body.Close()

	var posts []Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		log.Println("Error decoding posts:", err)
		return
	}

	// Print the first post as a sample
	fmt.Printf("Fetched %d posts. First post title: %s\n", len(posts), posts[0].Title)
}

func fetchComments(wg *sync.WaitGroup) {
	defer wg.Done() // Decrements the counter when the goroutine completes
	resp, err := http.Get("https://jsonplaceholder.typicode.com/comments")
	if err != nil {
		log.Println("Error fetching comments:", err)
		return
	}
	defer resp.Body.Close()

	var comments []Comment
	if err := json.NewDecoder(resp.Body).Decode(&comments); err != nil {
		log.Println("Error decoding comments:", err)
		return
	}

	// Print the first comment as a sample
	fmt.Printf("Fetched %d comments. First comment name: %s\n", len(comments), comments[0].Name)
}

func fetchUsers(wg *sync.WaitGroup) {
	defer wg.Done() // Decrements the counter when the goroutine completes
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Println("Error fetching users:", err)
		return
	}
	defer resp.Body.Close()

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		log.Println("Error decoding users:", err)
		return
	}

	// Print the first user as a sample
	fmt.Printf("Fetched %d users. First user name: %s\n", len(users), users[0].Name)
}

func main() {
	var wg sync.WaitGroup

	// Increment the counter for each concurrent task
	wg.Add(3)

	// Launch goroutines to fetch data concurrently
	go fetchPosts(&wg)
	go fetchComments(&wg)
	go fetchUsers(&wg)

	// Wait for all goroutines to finish
	wg.Wait()

	// Indicate that all tasks are completed
	fmt.Println("All data fetched concurrently!")
}
