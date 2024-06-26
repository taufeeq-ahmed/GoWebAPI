package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PageCount     int    `json:"page_count"`
	PublishedYear int    `json:"published_year"`
	IsAvailable   bool   `json:"is_available"`
}

var books = []Book{
	{Id: "fbhjwqkfr34", Title: "The Stranger", Author: "Albert Camus", IsAvailable: true, PageCount: 324, PublishedYear: 1987},
	{Id: "jhbjwqkjsbv", Title: "Great Gatsby", Author: "Scott F Fitzgerald", IsAvailable: true, PageCount: 123, PublishedYear: 1999},
}

func resppondToPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}

func getBookById(c *gin.Context) {
	var id = c.Param("id")
	var foundBook Book

	for _, book := range books {
		if book.Id == id {
			foundBook = book
			break
		}
	}

	if foundBook.Id != "" {
		c.JSON(http.StatusOK, foundBook)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	}
}

func createBook(c *gin.Context) {
	var newBook Book
	err := c.ShouldBindJSON(&newBook)

	if err == nil {
		books = append(books, newBook)
		c.JSON(http.StatusCreated, newBook)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
func updateBookById(c *gin.Context) {
	var id = c.Param("id")

	var index = -1
	for i, book := range books {
		if book.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var updatedBook Book
	err := c.ShouldBindJSON(&updatedBook)

	if err == nil {
		books[index] = updatedBook
		c.JSON(http.StatusOK, updatedBook)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

}

func deleteBookById(c *gin.Context) {
	var id = c.Param("id")

	var index = -1
	for i, book := range books {
		if book.Id == id {
			index = i
			break
		}
	}

	if index != -1 {
		books = append(books[:index], books[index+1:]...)
		c.JSON(http.StatusOK, gin.H{"message": "Book is deleted"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	}
}

func main() {
	router := gin.Default()
	router.GET("/ping", resppondToPing)
	router.GET("/books", getAllBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", createBook)
	router.PUT("/books/:id", updateBookById)
	router.DELETE("/books/:id", deleteBookById)
	router.Run(":9000")
}
