package models

type Dictionary struct {
	tableName struct{} `sql:"dictionary"`

	ID       int64  `json:"id",sql:"id"`
	Property string `json:"property",sql:"property"`
	Value    int    `json:"value",sql:"value"`
	Data     string `json:"data",sql:"data"`
}
