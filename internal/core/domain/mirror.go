package domain

import "time"

type Mirror struct {
	Name      string      `json:"name"`
	Position  string      `json:"position"`
	Email     string      `json:"email"`
	Hours     string      `json:"hours"`
	Month     string      `json:"month"`
	Year      int         `json:"year"`
	Registers []time.Time `json:"registers"`
}
