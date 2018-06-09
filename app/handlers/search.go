package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-pg/pg"
	"github.com/kot13/vertigo/app/container"
	"github.com/kot13/vertigo/app/renderer"
	"github.com/kot13/vertigo/db/models"
)

const MaxCountResults = 10000

var props = []SearchProperty{
	{
		Key:       "priceMin",
		Type:      "int",
		Title:     "Цена от",
		Condition: "advert_index.price >= ?",
	},
	{
		Key:       "priceMax",
		Type:      "int",
		Title:     "Цена до",
		Condition: "advert_index.price <= ?",
	},
	{
		Key:       "vendor",
		Type:      "[]int",
		Title:     "Производитель",
		Condition: "advert_index.vendor in (?)",
	},
}

func Search(w http.ResponseWriter, r *http.Request) {
	var adverts []models.AdvertData
	q := container.GetDb().Model(&adverts).ColumnExpr("advert_data.*").Join("JOIN advert_index ON advert_index.id = advert_data.id")

	values := r.URL.Query()
	for _, prop := range props {
		value := values.Get(prop.Key)
		if value != "" {
			switch prop.Type {
			case "int":
				q.Where(prop.Condition, value)
			case "string":
				q.Where(prop.Condition, value)
			case "[]int":
				items := strings.Split(value, ",")
				q.Where(prop.Condition, pg.In(items))
			case "[]string":
				items := strings.Split(value, ",")
				q.Where(prop.Condition, pg.In(items))
			}
		}
	}

	lat := values.Get("location[lat]")
	lon := values.Get("location[lon]")
	radius := values.Get("location[radius]")
	if lat != "" && lon != "" && radius != "" {
		q.Where(fmt.Sprintf("ST_DWithin(advert_index.point::geography,ST_GeomFromEWKT('SRID=4326;POINT(%s %s)')::geography, %s)", lon, lat, radius))
	}

	err := q.Limit(MaxCountResults).Select()
	if err != nil {
		renderer.Error(err.Error(), w)
		return
	}

	renderer.Render(adverts, w)
	return
}
