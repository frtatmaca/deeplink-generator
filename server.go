package main

import (
	"log"
	"net/http"

	logInstance "github.com/deeplink/core/logger"
	"github.com/deeplink/handler"
	"github.com/deeplink/helper"
	"github.com/deeplink/repository"
	"github.com/deeplink/service"
	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	l              logInstance.ILogger   = logInstance.NewLogger()
	linkContext    helper.LinkContext    = helper.NewLinkContext(l)
	linkRepository repository.Repository = repository.NewRepository()
	linkService    service.LinkService   = service.NewLinkService(linkRepository, linkContext, l)
	linkHandler    handler.LinkHandler   = handler.NewLinkHandler(linkService)
)

func init() {
}

// @title Go Restful API with Swagger
// @version 1.0
// @description Simple swagger implementation in Go HTTP

// @contact.name Firat Atmaca
// @contact.url frtatmaca
// @contact.email frtatmaca@gmail.com

// @host localhost:8000
// @BasePath /
func main() {
	http.Handle("/test", http.HandlerFunc(linkHandler.Test))
	http.Handle("/urltodeeplink", http.HandlerFunc(linkHandler.UrlToDeepLink))
	http.Handle("/deeplinktourl", http.HandlerFunc(linkHandler.DeepLinkToUrl))

	// Swagger
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
