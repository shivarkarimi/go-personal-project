package main

import (
	"net/http"

	"github.com/shivarkarimi/go-personal-project/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
