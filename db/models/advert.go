package models

import (
	"time"
)

type AdvertProperty struct {
	Data  string `json:"data",sql:"data"`
	Value string `json:"value",sql:"value"`
}

type AdvertData struct {
	ID         int64                     `json:"id",sql:"id"`
	HexID      string                    `json:"hexId",sql:"hex_id"`
	UserID     int64                     `json:"userId",sql:"user_id"`
	Properties map[string]AdvertProperty `json:"properties",sql:"properties"`
	CreatedAt  time.Time                 `json:"createdAt",sql:"created_at"`
	UpdatedAt  time.Time                 `json:"updatedAt",sql:"updated_at"`
}
