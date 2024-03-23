package dto

type FindDailyRegistryHeaders struct {
	UserId uint64 `header:"user-id" binding:"required"`
}
