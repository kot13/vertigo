package main

import (
	"os"
	"strconv"

	"github.com/kot13/vertigo/handlers"
	"github.com/kot13/vertigo/restapi"
	"github.com/kot13/vertigo/restapi/operations"
	"github.com/kot13/vertigo/restapi/operations/advert"
	"github.com/kot13/vertigo/restapi/operations/support"
	"github.com/kot13/vertigo/version"

	log "github.com/Sirupsen/logrus"
	"github.com/go-openapi/loads"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		log.Fatal("Logger level is not set.")
	}

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("Error parse log level: %s\n", err)
	} else {
		log.SetLevel(level)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	iport, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal("Port is not valid.")
	}

	log.Infof("Starting the service. Commit: %s, build time: %s, release: %s", version.Commit, version.BuildTime, version.Release)

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatal(err)
	}

	api := operations.NewVertigoAPI(swaggerSpec)

	api.Logger = log.Infof

	api.SupportGetHealthCheckHandler = support.GetHealthCheckHandlerFunc(handlers.HealthCheck)
	api.AdvertGetAdvertHandler = advert.GetAdvertHandlerFunc(handlers.GetAdverts)
	api.AdvertGetAdvertIDHandler = advert.GetAdvertIDHandlerFunc(handlers.GetAdvert)
	api.AdvertPostAdvertHandler = advert.PostAdvertHandlerFunc(handlers.CreateAdvert)
	api.AdvertPatchAdvertIDHandler = advert.PatchAdvertIDHandlerFunc(handlers.UpdateAdvert)
	api.AdvertDeleteAdvertIDHandler = advert.DeleteAdvertIDHandlerFunc(handlers.DeleteAdvert)

	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = iport

	err = server.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
