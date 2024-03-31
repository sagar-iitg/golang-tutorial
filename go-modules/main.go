package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

func main() {
	println("Hello, world!")
	router := httprouter.New()
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logger, _ := zap.NewProduction()
	logger.Info("successfully performed http request")

}
