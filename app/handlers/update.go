package handlers

import (
	"net/http"

	"github.com/kot13/vertigo/app/container"
	"github.com/kot13/vertigo/app/renderer"
	"github.com/kot13/vertigo/db/models"
)

type UpdateResponse struct {
	Success bool `json:"success"`
}

func Update(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if ID == "" {
		renderer.Error("Param id is require", w)
		return
	}

	props := map[string]models.AdvertProperty{
		"price": {
			Data:  "10000",
			Value: "10000",
		},
	}

	_, err := container.GetDb().SetProperties(ID, props)
	if err != nil {
		renderer.Error(err.Error(), w)
		return
	}

	renderer.Render(UpdateResponse{
		Success: true,
	}, w)
}
