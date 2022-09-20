package models

type Campaign struct {
	ID   uint   `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
