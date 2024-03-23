package dto

type RegisterDay struct {
	Name      string   `json:"name"`
	Position  string   `json:"position"`
	Hours     string   `json:"hours"`
	Registers []string `json:"registers"`
}
