package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/kot13/vertigo/models"
	"github.com/kot13/vertigo/restapi/operations/advert"
)

func CreateAdvert(_ advert.PostAdvertParams) middleware.Responder {
	payload := models.PostAdvertOKBody{
		ID:     1,
		Title:  "dummy advert",
		Status: "active",
		Price:  1000,
	}

	return advert.NewPostAdvertOK().WithPayload(&payload)
}
