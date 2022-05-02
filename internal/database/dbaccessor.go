package database

import "context"

// DBAccessor interface for databases
type DBAccessor interface {
	FetchRows(ctx context.Context, tag string, result interface{}, query string, args ...interface{}) (int, error)
}

// MysqlNoop used for testing implements DBAccessor interface
type MysqlNoop struct{}

// FetchRows noop
func (sql *MysqlNoop) FetchRows(ctx context.Context, tag string, result interface{}, query string, args ...interface{}) (int, error) {
	return 0, nil
}
