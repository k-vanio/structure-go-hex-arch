package ports

type Db interface {
	Close()
	AddHistory(answer int32, operation string) error
}
