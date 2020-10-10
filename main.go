package main

import (
	"net/http"

	"github.com/NileshGule/golang-bootcamp/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
