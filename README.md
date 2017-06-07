# parse-go
Parse.com API Client Library

This package provides a client for Parse's REST API. So far, it supports most of the query operations provided by Parse's Javascript library, with a few exceptions (listed below under TODO).

###Installation

```
go get github.com/ihoukai/parse-go

```
###Documentation Full documentation is provided by godoc.org

###Usage:


```


package main

import (
    "fmt"
	"time"
    
    parse "github.com/ihoukai/parse-go"
)

func main() {
	parse.Initialize(&parse.ClientConfig{
		Schema:     parse.SchemeHttp,
		Host:       "192.168.0.122:1337",
		PathPrefix: "/parse",
		Version:    "",
		AppID:      "myAppId",
		RestKey:    "restAPIKey",
		MasterKey:  "myAppId", // master key is optional
	})
    
    user := parse.User{}
    q, err := parse.NewQuery(&user)
	if err != nil {
		panic(err)
	}
    q.EqualTo("email", "ihoukai@gmail.com")
    q.GreaterThan("numFollowers", 10).OrderBy("-createdAt") // API is chainable
    err := q.First()
    if err != nil {
        if pe, ok := err.(parse.ParseError); ok {
            fmt.Printf("Error querying parse: %d - %s\n", pe.Code(), pe.Message())
        }
    }
    
    fmt.Printf("Retrieved user with id: %s\n", u.Id)

	q2, _ := parse.NewQuery(&parse.User{})
	q2.GreaterThan("createdAt", time.Date(2014, 01, 01, 0, 0, 0, 0, time.UTC))

	rc := make(chan *parse.User)

	// .Each will retrieve all results for a query and send them to the provided channel
	// The iterator returned allows for early cancelation of the iteration process, and
	// stores any error that triggers early termination
	iterator, err := q2.Each(rc)
	for u := range rc{
		fmt.Printf("received user: %v\n", u)
		// Do something
		if err := process(u); err != nil {
			// Cancel if there was an error
			iterator.Cancel()
		}
	}

	// An error occurred - not all rows were processed
	if it.Error() != nil {
		panic(it.Error())
	}
}

```

###TODO

Missing query operations
Related to
Missing CRUD operations:
Update
Field ops (__op):
AddRelation
RemoveRelation
Roles
Background Jobs
Analytics
File upload/retrieval
Batch operations
