package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dashdeipayan/mongodbgo/routers"
)

func main() {
	fmt.Println("Server is starting")
	r := routers.Router()
	log.Fatal(http.ListenAndServe(":4300", r))
	fmt.Println("listening at 4300")
}
