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

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func scanJob(message *workers.Msg) {
	// do something with your message
	// message.Jid()
	// message.Args() is a wrapper around go-simplejson (http://godoc.org/github.com/bitly/go-simplejson)
	fmt.Println("SSH to server")
	fmt.Println("Zip code")
	fmt.Println("SCP file zip to scanner")
	fmt.Println("Scan dependency")
	fmt.Println("Parser output fossa")
	fmt.Println("Dispatch job search CVE elasticsearch + export result")
	//TODO test share volume 2 container
	data := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}

	dirApp := "resources/app-" + message.Jid()
	if err := os.Mkdir(dirApp, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(dirApp+"/test.json", file, 0644)

	workers.Enqueue("sca-scanner-parser", "Add", PayloadParser{
		message.Jid(),
	})
}

type scanMiddleware struct{}

func (r *scanMiddleware) Call(queue string, message *workers.Msg, next func() bool) (acknowledge bool) {
	// do something before each message is processed
	acknowledge = next()
	// do something after each message is processed
	return
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

	workers.Middleware.Append(&scanMiddleware{})

	workers.Process("sca-scanner", scanJob, 10)

	// Add a job to a queue
	workers.Enqueue("sca-scanner", "Add", Payload{
		"192.168.5.1", "22", "22", "/user/local",
	})

	// stats will be available at http://localhost:8080/stats
	go workers.StatsServer(8080)
	workers.Run()
}

type Payload struct {
	Ip     string `json:"ip"`
	Server string `json:"server"`
	Port   string `json:"port"`
	Path   string `json:"path"`
}

type PayloadParser struct {
	JobId string `json:"job_id"`
}
