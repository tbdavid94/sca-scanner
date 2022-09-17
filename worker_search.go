package main

import (
	"fmt"
	"github.com/jrallison/go-workers"
)

func searchJob(message *workers.Msg) {
	// do something with your message
	// message.Jid()
	// message.Args() is a wrapper around go-simplejson (http://godoc.org/github.com/bitly/go-simplejson)
	payload := message.Args()
	jobIdNeedSearch := payload.Get("job_id").MustString()
	fmt.Println("Search CVE for source app-" + jobIdNeedSearch)
	fmt.Println("Connect ES")
	fmt.Println("Search CVE")
	fmt.Println("Export result")
}

func main() {
	workers.Configure(map[string]string{
		"namespace": "asset-sca",
		"server":    "127.0.0.1:6379",
		"database":  "0",
		"pool":      "30",
		"process":   "1",
	})

	workers.Process("sca-scanner-parser", searchJob, 20)

	// stats will be available at http://localhost:8080/stats
	go workers.StatsServer(8081)
	workers.Run()
}
