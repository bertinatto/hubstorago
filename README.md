# hubstorago

Hubstorago is an experimental go library for interacting with the
Scrapinghub API.
The main goal at the moment is to have a high-quality library with a strict
subset of functionalities implemented.


## Install

```bash
go get github.com/bertinatto/hubstorago
```

## Example

```go
package main

import (
	"fmt"
	"log"
	"github.com/bertinatto/hubstorago"
)

func main() {
	// First, make a client. It's possible to specify an alternative
	// base URL (for Hubstorage developers).
	c := hubstorago.Client{
		// BaseUrl: "http://my_dev_hostname",
		AuthKey: "********************************"}

	// Get a list of jobs from a given project.
	data, err := c.JobQ("1111111")
	if err != nil {
		log.Fatal(err)
	}
	for _, job := range *data {
		fmt.Printf("Job %s is %s\n", job.Key, job.State)
	}

	// Get the items the spider(s) scraped. Items returns an interface{}, so you
	// to translate the data yourself.
	items, err := c.Items("1111111/8/2")
	if err != nil {
		log.Fatal(err)
	}
	for _, rawItem := range (*items).([]interface{}) {
		item := rawItem.(map[string]interface{})
		fmt.Println(item["_url"])
	}

	// Get a list of requests made in the given job or project.
	requests, err := c.Requests("1111111/8/2")
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range *requests {
		fmt.Printf("%s - %d ms\n", r.Url, r.Duration)
	}

	// Logs from the job.
	logs, err := c.Logs("1111111/8/2")
	if err != nil {
		log.Fatal(err)
	}
	for _, log := range *logs {
		fmt.Println(log.Message)
	}

	// Some spider may store data in Collections.
	key := "testKey"
	value := "testValue"
	c.SetCollectionsKey("1111111", "s", "test", key, value)
	d, err := c.GetCollectionsKey("1111111", "s", "test", key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Value for key %s is %s\n", key, (*d)[0].Value)
}
```
