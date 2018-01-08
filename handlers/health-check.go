package handlers

import (
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/kot13/vertigo/models"
	"github.com/kot13/vertigo/restapi/operations"
	"github.com/kot13/vertigo/version"
)

func HealthCheck(params operations.GetHealthCheckParams) middleware.Responder {
	payload := &models.HealthCheck{
		Status:    "ok",
		BuildTime: version.BuildTime,
		Commit:    version.Commit,
		Release:   version.Release,
		Branch:    version.Branch,
	}

	return operations.NewGetHealthCheckOK().WithPayload(payload)
}
