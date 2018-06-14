package handlers

import (
	"net/http"

	"github.com/kot13/vertigo/app/container"
	"github.com/kot13/vertigo/app/renderer"
	"github.com/kot13/vertigo/db/models"
)

type PublishResponse struct {
	Success bool `json:"success"`
}

func Publish(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if ID == "" {
		renderer.Error("Param id is require", w)
		return
	}

	advert, err := container.GetDb().GetAdvertById(ID)
	if err != nil {
		renderer.Error(err.Error(), w)
		return
	}

	_, err = container.GetDb().SetStatus(ID, models.StatusSendToModeration)
	if err != nil {
		renderer.Error(err.Error(), w)
		return
	}

	err = container.GetDb().AddToIndex(ID, advert.Properties)
	if err != nil {
		renderer.Error(err.Error(), w)
		return
	}

	renderer.Render(PublishResponse{
		Success: true,
	}, w)
}
