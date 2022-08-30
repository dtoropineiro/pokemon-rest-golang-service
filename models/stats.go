package models

type Stats struct {
	HP        int `json:"hp"`
	Attack    int `json:"attack"`
	Defense   int `json:"defense"`
	Spattack  int `json:"spattack"`
	Spdefense int `json:"spdefense"`
	Speed     int `json:"speed"`
}
