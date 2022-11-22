package query

import "fmt"

type SQLSelectBuilder struct {
	TableName string
	FieldList string
	Where     []string
}

//goland:noinspection SqlNoDataSourceInspection
var (
	defaultQuery = "SELECT %s FROM %s"
)

func (q SQLSelectBuilder) ForTable(name string) SQLSelectBuilder {
	q.TableName = name
	return q
}

func (q SQLSelectBuilder) JustField(field string) SQLSelectBuilder {
	q.FieldList = field
	return q
}

func (q SQLSelectBuilder) AddWhere(where []string) SQLSelectBuilder {
	q.Where = where
	return q
}

func (q SQLSelectBuilder) Build() string {
	fieldList := func() string {
		if q.FieldList == "" {
			return "*"
		}

		return q.FieldList
	}

	query := fmt.Sprintf(defaultQuery, fieldList(), q.TableName)

	var where string
	if len(q.Where) > 0 {
		where += "WHERE "
		for _, w := range q.Where {
			where += fmt.Sprintf(`%s `, w)
		}
	}
	query = fmt.Sprintf("%s %s", query, where)

	return query
}
