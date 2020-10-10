package main

import (
	"net/http"

	"github.com/NileshGule/golang-bootcamp/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	_ = http.ListenAndServe(":3000", nil)
}
