package main

import (
	"fmt"
	"net/http"

	"github.com/binary-soup/amber/core/nodes/web"
	"github.com/binary-soup/amber/src/user"
	"github.com/julienschmidt/httprouter"
)

func createRouter() *httprouter.Router {
	r := httprouter.New()

	r.POST("/user", web.NewHandler(user.CreateEndpoint()).Serve)

	return r
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: createRouter(),
	}

	fmt.Println("Server running on port " + server.Addr[1:])
	server.ListenAndServe()
}
