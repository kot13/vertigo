package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	//"github.com/kot13/vertigo/models"
	"github.com/kot13/vertigo/restapi/operations"
)

func DeleteAdvert(_ operations.DeleteAdvertIDParams) middleware.Responder {
	return operations.NewDeleteAdvertIDNoContent()
}
