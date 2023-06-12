package sql

import (
	"fmt"
	"strings"
)

type SQLWhereOperator string

const (
	AndOperator SQLWhereOperator = "AND"
	OrOperator  SQLWhereOperator = "OR"
)

type SQLWhereBuilder struct {
	whereQueries []string
	whereValues  []interface{}
}

func NewSQLWhereBuilder() *SQLWhereBuilder {
	return &SQLWhereBuilder{
		whereQueries: []string{},
		whereValues:  []interface{}{},
	}
}

func (b *SQLWhereBuilder) AddQuery(query string, value interface{}) {
	b.whereQueries = append(b.whereQueries, query)
	b.whereValues = append(b.whereValues, value)
}

func (b *SQLWhereBuilder) GetQuery(operator SQLWhereOperator) string {
	return strings.Join(b.whereQueries, fmt.Sprintf(" %s ", operator))
}

func (b *SQLWhereBuilder) GetQueryValues() []interface{} {
	return b.whereValues
}

func (b *SQLWhereBuilder) IsEmpty() bool {
	if len(b.whereQueries) == 0 {
		return true
	}
	if len(b.whereValues) == 0 {
		return true
	}

	return false
}
