package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct{
	ID string `json: "id"`
	Title string `json: "title"`
	Author string `json: "author"`
	Quantity int `json: "quantity"`

}

var books = []book{
	{ID: "1", Title: "The Alchemist", Author: "Paulo Coelho", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "The Little Prince", Author: "Antoine de Saint-Exupéry", Quantity: 6},
}

