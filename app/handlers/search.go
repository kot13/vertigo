package handlers

import (
	"net/http"

	"fmt"
	"github.com/kot13/vertigo/app/container"
	"github.com/kot13/vertigo/app/renderer"
	"github.com/kot13/vertigo/db/models"
)

const MaxCountResults = 10000

var props = []SearchProperty{
	{
		Key:      "priceMin",
		Type:     "int64",
		Title:    "Цена от",
		Column:   "price",
		Operator: ">=",
	},
	{
		Key:      "priceMax",
		Type:     "int64",
		Title:    "Цена до",
		Column:   "price",
		Operator: "<=",
	},
}

func Search(w http.ResponseWriter, r *http.Request) {
	var adverts []models.AdvertData
	q := container.GetDb().Model(&adverts).ColumnExpr("advert_data.*").Join("JOIN advert_index ON advert_index.id = advert_data.id")

	values := r.URL.Query()
	for _, prop := range props {
		value := values.Get(prop.Key)
		if value != "" {
			q.Where(fmt.Sprintf("advert_index.%s %s ?", prop.Column, prop.Operator), value)
		}
	}

	err := q.Limit(MaxCountResults).Select()
	if err != nil {
		renderer.Error(err.Error(), w)
		return
	}

	renderer.Render(adverts, w)
	return
}
