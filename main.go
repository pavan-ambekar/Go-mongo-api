package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pavan-ambekar/Go-mongo-api/router"
)

func main() {
	fmt.Println("Mongo API")
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":3000", router.Router()))
	fmt.Println("Listening at port 3000...")
}
