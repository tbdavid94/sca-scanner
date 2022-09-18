package main

import (
	"github.com/joho/godotenv"
	"github.com/jrallison/go-workers"
	"log"
	"os"
)

func main() {
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

	// Add a job to a queue
	workers.Enqueue("sca-scanner", "Add", PayloadDispatch{
		"test-1.1.1.1", "test-22", "test-22", "test-/user/local",
	})
}

type PayloadDispatch struct {
	Ip     string `json:"ip"`
	Server string `json:"server"`
	Port   string `json:"port"`
	Path   string `json:"path"`
}
