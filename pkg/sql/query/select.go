package query

import (
	"fmt"
)

type SQLSelectBuilder struct {
	TableName string
	Fields    string
	Wheres    string
	Joins     string
	Order     string
	Group     string
	Limit     string
	Offset    string
}

//goland:noinspection SqlNoDataSourceInspection
var (
	defaultQuery = "SELECT %s FROM %s"
)

// Table
// Add table name
func (q SQLSelectBuilder) Table(name string) SQLSelectBuilder {
	q.TableName = name

	return q
}

// Field
// Add specified field to load
func (q SQLSelectBuilder) Field(fields string) SQLSelectBuilder {
	q.Fields = fields

	return q
}

// OrderBy
// add transaction by to query
// @usage
// builder.OrderBy(string)
// builder.OrderBy("name DESC")
// builder.OrderBy("name ASC")
// example query:
//
//	SELECT * FROM customers ORDER BY last_name, first_name;
func (q SQLSelectBuilder) OrderBy(order string) SQLSelectBuilder {
	q.Order = fmt.Sprintf("ORDER BY %s", order)

	return q
}

// GroupBy
// add group by to query
// @usage
// builder.GroupBy(string)
// builder.GroupBy("name")
// example query:
//
//	SELECT type, avg(age) AS average_age FROM customers GROUP BY type;
func (q SQLSelectBuilder) GroupBy(group string) SQLSelectBuilder {
	q.Group = fmt.Sprintf("GROUP BY %s", group)

	return q
}

// HasLimit
// Limit query result
func (q SQLSelectBuilder) HasLimit(limit int) SQLSelectBuilder {
	q.Limit = fmt.Sprintf("LIMIT %d", limit)

	return q
}

// HasOffset
// add offset to query
func (q SQLSelectBuilder) HasOffset(offset int) SQLSelectBuilder {
	q.Offset = fmt.Sprintf("OFFSET %d", offset)

	return q
}

// Build
// append the query string
func (q SQLSelectBuilder) Build() string {
	if q.Fields == "" {
		q.Fields = "*"
	}

	query := fmt.Sprintf(defaultQuery, q.Fields, q.TableName)

	query = fmt.Sprintf(
		"%s %s %s %s %s %s %s",
		query,
		q.Wheres,
		q.Joins,
		q.Group,
		q.Order,
		q.Limit,
		q.Offset,
	)

	return query
}
