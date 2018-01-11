package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/kot13/vertigo/models"
	"github.com/kot13/vertigo/restapi/operations"
)

func GetAdverts(_ operations.GetAdvertParams) middleware.Responder {
	payload := []*models.Advert{
		{
			ID:     1,
			Title:  "dummy advert",
			Status: "active",
			Price:  1000,
		},
	}

	return operations.NewGetAdvertOK().WithPayload(payload)
}
