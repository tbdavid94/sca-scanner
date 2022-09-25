package main

import (
	"context"
	"fmt"
	config2 "github.com/olivere/elastic/config"
	"gopkg.in/olivere/elastic.v6"
	"gopkg.in/olivere/elastic.v6/config"
)

func main() {
	configString, _ := config.Parse("http://127.0.0.1:9200/index?sniff=false")
	client, err := elastic.NewClientFromConfig((*config2.Config)(configString))
	if err != nil {
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	exists, err := client.IndexExists("cve-index").Do(context.Background())
	if err != nil {
		// Handle error
	}
	if !exists {
		fmt.Println("404")
	}
}
