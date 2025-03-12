package main

import "test/routes"

func main() {

	r := routes.InitRouter()
	r.Run(":8080")

}
