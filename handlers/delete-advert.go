package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	//"github.com/kot13/vertigo/models"
	"github.com/kot13/vertigo/restapi/operations/advert"
)

func DeleteAdvert(_ advert.DeleteAdvertIDParams) middleware.Responder {
	return advert.NewDeleteAdvertIDNoContent()
}
