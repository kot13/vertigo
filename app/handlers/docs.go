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

type data struct {
	Advert     []Property
	Properties []Property
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
	}

	t.Execute(w, d)
}
