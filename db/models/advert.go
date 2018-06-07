package models

import (
	"time"
)

const (
	StatusNew              = 1
	StatusSendToModeration = 2
	StatusOnModeration     = 3
	StatusPublished        = 4
	StatusUnPublished      = 5
	StatusNotValid         = 6
)

type AdvertProperty struct {
	Data  string `json:"data",sql:"data"`
	Value string `json:"value",sql:"value"`
}

type AdvertData struct {
	tableName struct{} `sql:"advert_data",alias:"advert_data"`

	ID         int64                     `json:"id",sql:"id"`
	UserID     int64                     `json:"userId",sql:"user_id"`
	Properties map[string]AdvertProperty `json:"properties",sql:"properties"`
	CreatedAt  time.Time                 `json:"createdAt",sql:"created_at"`
	UpdatedAt  *time.Time                `json:"updatedAt",sql:"updated_at"`
	Status     int8                      `json:"status",sql:"status"`
}

type AdvertIndex struct {
	tableName struct{} `sql:"advert_index",alias:"advert_index"`

	Price int64 `sql:"price"`
}
