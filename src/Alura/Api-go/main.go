package main

import (
	"fmt"
	"net/http"

	"api-web/routes"

)

func main() {
	routes.CarregaRotas()

	fmt.Println("Starting server on port 8000.")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("ERROR -> Fail to start server")
		panic(err.Error())
	}
}
