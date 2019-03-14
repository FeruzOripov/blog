package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Blog struct for service.
type Blog struct {
	ID           int        `json:"id"`
	Title        string     `json:"title"`
	Body         string     `json:"body"`
	CreatedDay   int        `json:"createdDay"`
	CreatedMonth time.Month `json:"createdMonth`
	CreatedYear  int        `json:"createdYear"`
}

// Initializing id. Each time creating new post it is incremented
var id int

// Initiliazing some posts.
var posts = []Blog{
	Blog{id, "Hello, World!", "This is test", time.Now().Day(), time.Now().Month(), time.Now().Year()},
}

func main() {
	r := gin.Default()
	// Defining a group named /api
	api := r.Group("/api")
	{
		// Test url that responses with Hello, World!
		api.GET("/", sayHello)
		// Lists all posts
		api.GET("/posts", getPosts)
		// Get a post with given id
		api.GET("/posts/:id", getPost)
		// Post a new post
		api.POST("/posts", postPost)
		// Update a post with given id
		api.POST("/posts/:id", updatePost)
		// Delete a post with a given id
		api.DELETE("posts/:id", deletePost)
	}
	r.Run()
}

// responses with Hello, World
func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

// List all posts
func getPosts(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, posts)
}

// Get a post with given id
func getPost(c *gin.Context) {
	if postID, err := strconv.Atoi(c.Param("id")); err == nil {
		for _, post := range posts {
			if post.ID == postID {
				c.Header("Content-Type", "application/json")
				c.JSON(http.StatusOK, post)
				return
			}
		}
		c.JSON(404, "Not found")
	}
}

// Create a new post
func postPost(c *gin.Context) {
	var post Blog
	c.Bind(&post)
	id++ // Increment an id for new post
	post.ID = id
	post.CreatedDay = time.Now().Day()
	post.CreatedMonth = time.Now().Month()
	post.CreatedYear = time.Now().Year()
	posts = append(posts, post)
	c.JSON(201, post)
}

// Update a post
func updatePost(c *gin.Context) {
	var post Blog
	c.Bind(&post)
	if postID, err := strconv.Atoi(c.Param("id")); err == nil {
		for _, p := range posts {
			if p.ID == postID {
				p.Title = post.Title
				p.Body = post.Body
				c.Header("Content-Type", "application/json")
				c.JSON(http.StatusOK, p)
				return
			}
		}
		c.JSON(404, "Not found")
	}
}

// Delete a post
func deletePost(c *gin.Context) {
	if postID, err := strconv.Atoi(c.Param("id")); err == nil {
		for i, post := range posts {
			if post.ID == postID {
				posts = removeIndex(posts, i)
				c.Header("Content-Type", "application/json")
				c.JSON(http.StatusOK, posts)
				return
			}
		}
		c.JSON(404, "Not found")
	}
}

// Function for removing element from array.
// Because our posts are in array.
func removeIndex(s []Blog, index int) []Blog {
	return append(s[:index], s[index+1:]...)
}
