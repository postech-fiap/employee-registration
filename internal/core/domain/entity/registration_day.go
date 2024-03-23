package entity

type DailyRegistry struct {
	Name          string   `json:"name"`
	Position      string   `json:"position"`
	Hours         string   `json:"hours"`
	DailyRegistry []string `json:"daily_registry"`
}
