package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/jrallison/go-workers"
	"io/ioutil"
	"log"
	"os"
)

func searchJob(message *workers.Msg) {
	// do something with your message
	// message.Jid()
	// message.Args() is a wrapper around go-simplejson (http://godoc.org/github.com/bitly/go-simplejson)
	payload := message.Args()
	jobIdNeedSearch := payload.Get("job_id").MustString()
	fmt.Println("Search CVE for source app-" + jobIdNeedSearch)

	jsonFile, err := os.Open("resources/app-" + jobIdNeedSearch + "/test.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened test.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result["MonthlySalary"])

	fmt.Println("Connect ES")
	fmt.Println("Search CVE")
	fmt.Println("Export result")
}

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

	workers.Process("sca-scanner-parser", searchJob, 20)

	// stats will be available at http://localhost:8081/stats
	go workers.StatsServer(8081)
	workers.Run()
}
