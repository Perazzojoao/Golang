package routes

import (
	"net/http"

	"api-web/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)

}
