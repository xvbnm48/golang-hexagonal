package ports

type DBPort interface {
	ClosedDbConnection()
	AddToHistory(answer int32, operation string) error
}
