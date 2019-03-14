package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Blog struct {
	ID           int        `json:"id"`
	Title        string     `json:"title"`
	Body         string     `json:"body"`
	CreatedDay   int        `json:"createdDay"`
	CreatedMonth time.Month `json:"createdMonth`
	CreatedYear  int        `json:"createdYear"`
}

var id int

var posts = []Blog{
	Blog{id, "Hello, World!", "This is test", time.Now().Day(), time.Now().Month(), time.Now().Year()},
}

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/", sayHello)
	}
	api.GET("/posts", getPosts)
	api.GET("/posts/:id", getPost)
	api.POST("/posts", postPost)
	api.POST("/posts/:id", updatePost)
	api.DELETE("posts/:id", deletePost)
	r.Run()
}

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func getPosts(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, posts)
}

func getPost(c *gin.Context) {
	if postID, err := strconv.Atoi(c.Param("id")); err == nil {
		for _, post := range posts {
			if post.ID == postID {
				//c.Header("Content-Type", "application/json")
				c.JSON(http.StatusOK, post)
				return
			}
		}
		c.JSON(404, "Not found")
	}
}

func postPost(c *gin.Context) {
	var post Blog
	c.Bind(&post)
	id++
	post.ID = id
	post.CreatedDay = time.Now().Day()
	post.CreatedMonth = time.Now().Month()
	post.CreatedYear = time.Now().Year()
	posts = append(posts, post)
	c.JSON(201, post)
}

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

func deletePost(c *gin.Context) {
	if postID, err := strconv.Atoi(c.Param("id")); err == nil {
		for i, post := range posts {
			if post.ID == postID {
				posts = RemoveIndex(posts, i)
				c.Header("Content-Type", "application/json")
				c.JSON(http.StatusOK, posts)
				return
			}
		}
		c.JSON(404, "Not found")
	}
}

func RemoveIndex(s []Blog, index int) []Blog {
	return append(s[:index], s[index+1:]...)
}
