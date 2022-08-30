package models

type Type struct {
	ID   string   `json:"id"`
	Name string   `json:"name,omitempty" validate:"required"`
	Img  string   `json:"img"`
	Type []string `json:"type"`
}
