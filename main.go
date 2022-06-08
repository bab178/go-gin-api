package main

import (
	"errors"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct {
	Id		string	`json:"id"`
	Title 	string	`json:"title"`
	Author	string	`json:"author"`
	Quantity int	`json:"quantity"`
}

var books = []book {
	{Id: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{Id: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{Id: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

// Allow any CORS for local development
func CORSMiddleware() gin.HandlerFunc {
	fmt.Print("here")
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := findBook(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func findBook(id string) (*book, error) {
	for i, b := range books {
		if b.Id == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, books)
}

func tryUpdateBookQuantity(id string, num int) error {
	book, err := findBook(id)
	
	if err != nil {
		return err 
	}

	if num >= 0 || book.Quantity > 0 {
		book.Quantity += num
		return nil
	} else {
		return fmt.Errorf("no books with id %s available to checkout", id)
	}
}

func checkoutBook(c *gin.Context) {
	id := c.Param("id")
	err := tryUpdateBookQuantity(id, -1)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, books)
}

func returnBook(c *gin.Context) {
	id := c.Param("id")
	err := tryUpdateBookQuantity(id, 1)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	// Allow any CORS for local development
	router.Use(CORSMiddleware())

	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", createBook)
	router.POST("/books/checkout/:id", checkoutBook)
	router.POST("/books/return/:id", returnBook)

	router.Run("localhost:8080")
}