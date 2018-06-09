package handlers

import (
	"fmt"
	"net/http"

	"encoding/json"
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
	{
		Key:      "location",
		Type:     "location",
		Title:    "Область поиска",
		Column:   "point",
		Operator: "=",
	},
}

type Location struct {
	Lat    string `json:"lat"`
	Lon    string `json:"lon"`
	Radius string `json:"radius"`
}

func Search(w http.ResponseWriter, r *http.Request) {
	var adverts []models.AdvertData
	q := container.GetDb().Model(&adverts).ColumnExpr("advert_data.*").Join("JOIN advert_index ON advert_index.id = advert_data.id")

	values := r.URL.Query()
	for _, prop := range props {
		value := values.Get(prop.Key)
		if value != "" {
			switch prop.Type {
			case "location":
				var loc Location
				err := json.Unmarshal([]byte(value), &loc)
				if err != nil {
					renderer.Error(err.Error(), w)
					return
				}

				q.Where(
					fmt.Sprintf(`ST_DWithin(advert_index.point::geography,ST_GeomFromEWKT('SRID=4326;POINT(%s %s)')::geography, %s)`,
						loc.Lon, loc.Lat, loc.Radius))
			default:
				q.Where(fmt.Sprintf("advert_index.%s %s ?", prop.Column, prop.Operator), value)
			}
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
