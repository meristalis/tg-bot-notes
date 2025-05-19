package entity

type Tag struct {
	ID   int    `json:"id" db:"id" example:"1"`
	Name string `json:"name" db:"name" example:"Important"`
}
