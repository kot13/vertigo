package models

import (
	"time"
)

type AdvertProperties struct {
}

type AdvertData struct {
	ID         int64            `json:"id",sql:"id"`
	HexID      string           `json:"hexId",sql:"hex_id"`
	UserID     int64            `json:"userId",sql:"user_id"`
	Properties AdvertProperties `json:"properties",sql:"properties"`
	CreatedAt  time.Time        `json:"createdAt",sql:"created_at"`
	UpdatedAt  time.Time        `json:"updatedAt",sql:"updated_at"`
}
