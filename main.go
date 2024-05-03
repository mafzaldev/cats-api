package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Cat struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

var cats = []Cat{
	{ID: "1", Name: "Whiskers", Image: "https://cataas.com/cat/FmTkbYVy0FSi4QBl"},
	{ID: "2", Name: "Fluffy", Image: "https://cataas.com/cat/IAYCOahG15FFb3Qi"},
	{ID: "3", Name: "Mittens", Image: "https://cataas.com/cat/ughE5w0OlXauQ81k"},
	{ID: "4", Name: "Smokey", Image: "https://cataas.com/cat/8hJ4kU7EtjB8Y203"},
	{ID: "5", Name: "Tigger", Image: "https://cataas.com/cat/VJi5njDrqdQxpDvx"},
	{ID: "6", Name: "Oreo", Image: "https://cataas.com/cat/2TZUgzYXLM9SzFmX"},
}

func index(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte("API is running 🚀..."))
}

func getCats(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, cats)
}

func getCatByID(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, cat := range cats {
		if cat.ID == id {
			ctx.IndentedJSON(http.StatusOK, cat)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Cat not found!"})
}

func getRandomCat(ctx *gin.Context) {
	min, max := 0, len(cats)-1
	randomIndex := rand.Intn(max-min+1) + min
	ctx.IndentedJSON(http.StatusOK, cats[randomIndex])
}

func main() {

	router := gin.Default()
	router.GET("/", index)
	router.GET("/cats", getCats)
	router.GET("/cats/:id", getCatByID)
	router.GET("/cats/random", getRandomCat)

	router.Run("localhost:3000")
}
