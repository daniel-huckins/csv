package csv

// Config contains values about the application
type Config interface {
	DBName() string
}
