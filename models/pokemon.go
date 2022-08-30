package models

type Pokemon struct {
	ID    string   `json:"id"`
	Name  string   `json:"name,omitempty" validate:"required"`
	Img   string   `json:"img"`
	Type  []string `json:"type"`
	Stats Stats    `json:"stats"`
}
