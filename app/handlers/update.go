package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/kot13/vertigo/app/container"
	"github.com/kot13/vertigo/app/renderer"
	"github.com/kot13/vertigo/db/models"
)

type UpdateRequest struct {
	Price *int64  `json:"price,omitempty"`
	Lat   *string `json:"lat,omitempty"`
	Lon   *string `json:"lon,omitempty"`
}

type UpdateResponse struct {
	Success bool `json:"success"`
}

func Update(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if ID == "" {
		renderer.Error("Param id is require", w)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		renderer.Error(err.Error(), w)
		return
	}

	var req UpdateRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		renderer.Error(err.Error(), w)
		return
	}

	props := models.AdvertProperties{
		Price: req.Price,
		Lat:   req.Lat,
		Lon:   req.Lon,
	}

	_, err = container.GetDb().SetProperties(ID, props)
	if err != nil {
		renderer.Error(err.Error(), w)
		return
	}

	renderer.Render(UpdateResponse{
		Success: true,
	}, w)
}
