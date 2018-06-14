package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/kot13/vertigo/app/container"
	"github.com/kot13/vertigo/app/renderer"
	"github.com/kot13/vertigo/db/models"
)

type CreateRequest struct {
	UserID int64 `json:"userId"`
}

type CreateResponse struct {
	ID int64 `json:"id"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		renderer.Error(err.Error(), w)
		return
	}

	var req CreateRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		renderer.Error(err.Error(), w)
		return
	}

	a := models.AdvertData{
		UserID: req.UserID,
		Status: models.StatusNew,
	}
	err = container.GetDb().Insert(&a)
	if err != nil {
		renderer.Error(err.Error(), w)
		return
	}

	renderer.Render(CreateResponse{
		ID: a.ID,
	}, w)
}
