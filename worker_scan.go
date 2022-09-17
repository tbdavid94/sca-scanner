package main

import (
	"fmt"
	"github.com/jrallison/go-workers"
)

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
	workers.Configure(map[string]string{
		"namespace": "asset-sca",
		"server":    "127.0.0.1:6379",
		"database":  "0",
		"pool":      "30",
		"process":   "1",
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
