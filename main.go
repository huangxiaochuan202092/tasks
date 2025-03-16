package main

import (
	"fmt"
	"log"
	"proapp/routes"
)

func main() {
	r := routes.InitRouter()

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}
