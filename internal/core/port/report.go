package port

type ReportUseCase interface {
	GenarateMirrorReport(userId int) error
}
