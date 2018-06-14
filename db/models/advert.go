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

type AdvertData struct {
	tableName struct{} `sql:"advert_data",alias:"advert_data"`

	ID         int64            `json:"id",sql:"id"`
	UserID     int64            `json:"userId",sql:"user_id"`
	Properties AdvertProperties `json:"properties",sql:"properties"`
	CreatedAt  time.Time        `json:"createdAt",sql:"created_at"`
	UpdatedAt  *time.Time       `json:"updatedAt",sql:"updated_at"`
	Status     int8             `json:"status",sql:"status"`
}

type AdvertProperties struct {
	Price  *int64  `json:"price,omitempty"`
	Vendor *int64  `json:"vendor,omitempty"`
	Lat    *string `json:"lat,omitempty"`
	Lon    *string `json:"lon,omitempty"`
}

type AdvertIndex struct {
	tableName struct{} `sql:"advert_index",alias:"advert_index"`

	Id     int64  `sql:"id"`
	Price  int64  `sql:"price"`
	Vendor int64  `sql:"vendor"`
	Point  string `sql:"point"`
}
