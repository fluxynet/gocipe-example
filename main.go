package main

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/fluxynet/gocipe-example/app"
	"github.com/fluxynet/gocipe-example/models"
	"github.com/gorilla/mux"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

//go:generate rice embed-go

func main() {
	config := app.Bootstrap()
	models.Init(config.DB)

	g := grpc.NewServer()
	ws := grpcweb.WrapServer(g)

	router := mux.NewRouter()
	router.PathPrefix("/api").Handler(http.StripPrefix("/api", ws))

	if app.Env == app.EnvironmentDev {
		router.PathPrefix("/").Handler(http.FileServer(http.Dir("web/dist")))
	} else {
		assetsHandler := http.FileServer(rice.MustFindBox("web/dist").HTTPBox())
		router.PathPrefix("/").Handler(assetsHandler)
	}

	http.ListenAndServe(":"+config.HTTPPort, router)
}
