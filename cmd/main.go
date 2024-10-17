package main

import (
	_ "Pinterest/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/swaggo/files"
	//"github.com/swaggo/gin-swagger"
	"os"
)

// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} HTTPError
// @Router /users/{id} [get]
func getusers(c *gin.Context) {
	c.JSON(200, gin.H{
		"email": "john@example.com",
	})
}

func main() {
	path, exists := os.LookupEnv("PATH")

	if exists {
		// Print the value of the environment variable
		fmt.Print(path)
	}
}
