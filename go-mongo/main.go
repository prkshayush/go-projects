package main

import (
	"fmt"
	"go-mongo/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("MongoDBAPI")
	r := router.Router()
	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":3400", r))
	fmt.Println("Listening at port 3400...")
}
