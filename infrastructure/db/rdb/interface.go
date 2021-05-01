package rdb

type SQLHandler interface {
	Query(interface{}, string, ...interface{}) error
}
