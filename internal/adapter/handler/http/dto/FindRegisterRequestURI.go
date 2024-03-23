package dto

type FindAllRegisterRequestURI struct {
	UserId uint64 `uri:"id" binding:"required,gt=0"`
}
