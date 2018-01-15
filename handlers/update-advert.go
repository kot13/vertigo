package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/kot13/vertigo/models"
	"github.com/kot13/vertigo/restapi/operations/advert"
)

func UpdateAdvert(_ advert.PatchAdvertIDParams) middleware.Responder {
	payload := models.PatchAdvertIDOKBody{
		ID:     1,
		Title:  "dummy advert",
		Status: "active",
		Price:  1000,
	}

	return advert.NewPatchAdvertIDOK().WithPayload(&payload)
}
