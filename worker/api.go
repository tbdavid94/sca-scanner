package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/jrallison/go-workers"
	"log"
	"net/http"
	"os"
)

type scan struct {
	IP       string `json:"ip" binding:"required"`
	Port     string `json:"port" binding:"required"`
	Path     string `json:"path" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.POST("/scan/add", addScan)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	workers.Configure(map[string]string{
		"namespace": os.Getenv("REDIS_NAMESPACE"),
		"server":    os.Getenv("REDIS_SERVER"),
		"database":  os.Getenv("REDIS_DB"),
		"pool":      os.Getenv("REDIS_POOL"),
		"process":   os.Getenv("REDIS_PROCESS"),
	})

	router.Run("localhost:8083")
}

func ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "pong",
	})
}

func addScan(c *gin.Context) {
	var newScan scan

	// Call BindJSON to bind the received JSON to
	// newScan.
	if err := c.ShouldBindJSON(&newScan); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	fmt.Println(newScan)
	//c.JSON(http.StatusAccepted, &newScan)
	// Add a job to a queue
	wk, err := workers.Enqueue("sca-scanner", "Add", newScan)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dispatch job success", wk)
	c.IndentedJSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "dispatch job success",
		"data":    wk,
	})

}
