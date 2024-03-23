package dto

type NewRegisterMessage struct {
	ID string `json:"user_id" validate:"required,gt=0"`
}
