package controllers

/*
	Module to handle the api request
*/

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TMaTTeO99/UrlHealthMonitor/API/middleware"
	"github.com/TMaTTeO99/UrlHealthMonitor/API/service"
	"github.com/TMaTTeO99/UrlHealthMonitor/config"
)

func StartWebServer(config *config.ConfigData) {

	mux := http.NewServeMux()

	urlService := &service.UrlService{
		Config: config,
		Client: &http.Client{},
	}
	mux.HandleFunc("GET /url/verified-url/{urlValue}", urlService.SearchHandling)
	mux.HandleFunc("POST /url/verified-url/:findInfo", urlService.AnalizeHandling)
	handler := middleware.ApplyCorse(mux)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")), handler))

}
