package query

import "fmt"

// ==================================================================================
// ---------------------------------- JOIN TABLE ----------------------------------//
// ==================================================================================

// Join
// add join between two tables
// @usage
// join(string)
// builder.Join("cities c on t.id = c.id")
// builder.Join([]string{"cities c on t.id = c.id", "province p on c.id = p.id")
func combineJoin(joinValue interface{}, typeJoin string) string {
	if j, ok := joinValue.(string); ok {
		return fmt.Sprintf("%s %s ", typeJoin, j)
	}

	if j, ok := joinValue.([]string); ok {
		var joins string

		joins += fmt.Sprintf("%s ", typeJoin)

		for i, data := range j {
			if i != 0 {
				joins += fmt.Sprintf("%s ", typeJoin)
			}

			joins += fmt.Sprintf("%s ", data)
		}

		return joins
	}

	return ""
}

// Join
// basic join query
func (q SQLSelectBuilder) Join(join interface{}) SQLSelectBuilder {
	q.Joins += combineJoin(join, "JOIN")

	return q
}

// InnerJoin
// returns rows when there is a match in both tables.
func (q SQLSelectBuilder) InnerJoin(join interface{}) SQLSelectBuilder {
	q.Joins += combineJoin(join, "INNER JOIN")

	return q
}

// CrossJoin
// A CROSS JOIN matches every row of the first table with every row of the second table.
// If the input tables have x and y columns, respectively, the resulting table will have x+y columns.
func (q SQLSelectBuilder) CrossJoin(join interface{}) SQLSelectBuilder {
	q.Joins += combineJoin(join, "CROSS JOIN")

	return q
}

// LeftOuterJoin
// for each row in table T1 that does not satisfy
// the join condition with any row in table T2,
// a joined row is added with null values in columns of T2.
// Thus, the joined table always has at least one row for each row in T1.
func (q SQLSelectBuilder) LeftOuterJoin(join interface{}) SQLSelectBuilder {
	q.Joins += combineJoin(join, "LEFT OUTER JOIN")

	return q
}

// RightOuterJoin
// for each row in table T2 that does not satisfy
// the join condition with any row in table T1,
// a joined row is added with null values in columns of T1.
// This is the converse of a left join; the result table will always have a row for each row in T2.
func (q SQLSelectBuilder) RightOuterJoin(join interface{}) SQLSelectBuilder {
	q.Joins += combineJoin(join, "RIGHT OUTER JOIN")

	return q
}

// FullOuterJoin
// for each row in table T1 that does not satisfy the join condition with any row in table T2,
// a joined row is added with null values in columns of T2.
// In addition, for each row of T2 that does not satisfy the join condition with any row in T1,
// a joined row with null values in the columns of T1 is added.
func (q SQLSelectBuilder) FullOuterJoin(join interface{}) SQLSelectBuilder {
	q.Joins += combineJoin(join, "FULL OUTER JOIN")

	return q
}
