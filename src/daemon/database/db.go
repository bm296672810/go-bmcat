package database

type Db interface {
	Connect() bool
	IsConnected() bool
	Tables() []string
}
