package entity

import "time"

type DailyRegistry struct {
	Name          string      `json:"name"`
	Position      string      `json:"position"`
	Hours         string      `json:"hours"`
	DailyRegistry []time.Time `json:"daily_registry"`
}
