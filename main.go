package main

import (
	"github.com/go_restapi/server"
)

func main() {

	// Endpoint 8000
	srv := server.New(":8080")

	// Handle error
	error := srv.ListenAndServe()

	if error != nil {
		panic(error)
	}

}
