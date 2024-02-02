package main

import (
	"fmt"
	"net/http"

	"api-web/routes"

)

func main() {
	fmt.Println("Servidor online!")

	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
