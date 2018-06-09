package handlers

import (
	"html/template"
	"net/http"
)

type Property struct {
	Key        string
	Type       string
	Title      string
	IsRequired bool
	IsIndex    bool
}

type SearchProperty struct {
	Key       string
	Type      string
	Title     string
	Example   string
	Condition string
	Column    string
	Operator  string
}

type data struct {
	Advert         []Property
	Properties     []Property
	SearchProperty []SearchProperty
}

func Docs(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/docs.html")

	var d data

	d.Advert = []Property{
		{
			Key:        "id",
			Type:       "int64",
			Title:      "Уникальный идентификатор",
			IsRequired: true,
		},
		{
			Key:        "userId",
			Type:       "int64",
			Title:      "ID пользователя",
			IsRequired: true,
		},
		{
			Key:        "properties",
			Type:       "map[string]Object",
			Title:      "Map свойств объявления вида ключ-значение",
			IsRequired: true,
		},
		{
			Key:        "createdAt",
			Type:       "string",
			Title:      "Дата создания",
			IsRequired: true,
		},
		{
			Key:        "updatedAt",
			Type:       "string",
			Title:      "Дата изменения",
			IsRequired: false,
		},
		{
			Key:        "status",
			Type:       "int8",
			Title:      "Статус объявления",
			IsRequired: true,
		},
	}

	d.Properties = []Property{
		{
			Key:        "price",
			Type:       "int64",
			Title:      "Цена товара в рублях",
			IsRequired: true,
			IsIndex:    true,
		},
		{
			Key:        "lat",
			Type:       "string",
			Title:      "Широта",
			IsRequired: true,
			IsIndex:    true,
		},
		{
			Key:        "lon",
			Type:       "string",
			Title:      "Долгота",
			IsRequired: true,
			IsIndex:    true,
		},
	}

	d.SearchProperty = []SearchProperty{
		{
			Key:     "priceMin",
			Type:    "int64",
			Title:   "Цена от",
			Example: "1",
		},
		{
			Key:     "priceMax",
			Type:    "int64",
			Title:   "Цена до",
			Example: "100000",
		},
		{
			Key:     "location[lat]",
			Type:    "string",
			Title:   "Область поиска (широта)",
			Example: "60.009354",
		},
		{
			Key:     "location[lon]",
			Type:    "string",
			Title:   "Область поиска (долгота)",
			Example: "30.326351",
		},
		{
			Key:     "location[radius]",
			Type:    "string",
			Title:   "Область поиска (радиус)",
			Example: "1000",
		},
	}

	t.Execute(w, d)
}
