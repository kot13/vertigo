package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/kot13/vertigo/models"
	"github.com/kot13/vertigo/restapi/operations"
	"github.com/kot13/vertigo/version"
)

func HealthCheck(_ operations.GetHealthCheckParams) middleware.Responder {
	payload := &models.HealthCheck{
		Status:    models.HealthCheckStatusOk,
		BuildTime: version.BuildTime,
		Commit:    version.Commit,
		Release:   version.Release,
		Branch:    version.Branch,
	}

	return operations.NewGetHealthCheckOK().WithPayload(payload)
}
