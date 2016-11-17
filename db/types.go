package db

import (
	"reflect"
)

// table is a relational data table
type table struct {
	ID   int64
	Name string
}

// column
type column struct {
	TableID int64
	Type    reflect.Type
}
