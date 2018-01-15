package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/kot13/vertigo/models"
	"github.com/kot13/vertigo/restapi/operations/advert"
)

func GetAdverts(_ advert.GetAdvertParams) middleware.Responder {
	payload := models.GetAdvertOKBody{
		{
			ID:     1,
			Title:  "dummy advert",
			Status: "active",
			Price:  1000,
		},
	}

	return advert.NewGetAdvertOK().WithPayload(payload)
}
