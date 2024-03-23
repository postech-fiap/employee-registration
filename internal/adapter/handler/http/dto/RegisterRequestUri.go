package dto

type RegisterRequestURI struct {
	ID int64 `uri:"employee_id" binding:"required,gt=0"`
}
