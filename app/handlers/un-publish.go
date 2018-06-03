package handlers

import (
	"net/http"

	"github.com/kot13/vertigo/app/container"
	"github.com/kot13/vertigo/app/renderer"
	"github.com/kot13/vertigo/db/models"
)

type UnPublishResponse struct {
	Success bool `json:"success"`
}

func UnPublish(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if ID == "" {
		renderer.Error("Param id is require", w)
		return
	}

	_, err := container.GetDb().SetStatus(ID, models.StatusUnPublished)
	if err != nil {
		renderer.Error(err.Error(), w)
		return
	}

	renderer.Render(UnPublishResponse{
		Success: true,
	}, w)
}
