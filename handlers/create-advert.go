package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/kot13/vertigo/models"
	"github.com/kot13/vertigo/restapi/operations"
)

func CreateAdvert(_ operations.PostAdvertParams) middleware.Responder {
	payload := models.Advert{
		ID:     1,
		Title:  "dummy advert",
		Status: "active",
		Price:  1000,
	}

	return operations.NewPostAdvertOK().WithPayload(&payload)
}
