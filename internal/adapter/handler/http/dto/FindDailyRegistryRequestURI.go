package dto

type FindDailyRegistryRequestURI struct {
	UserId uint64 `uri:"id" binding:"required,gt=0"`
}
