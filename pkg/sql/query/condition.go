package query

// ==================================================================================
// --------------------------------- WHERE CLAUSE ---------------------------------//
// ==================================================================================

// TODO
func combineConditions() string {
	return ""
}

// Where
// add where clause with specified operator default operator (and)
// e.g:
// where id = 1
// where id = 1 and name = 'json'
// first_name IN ('Ann','Anne','Annie')
// first_name LIKE 'Ann%'
// first_name LIKE 'A%' AND LENGTH(first_name) BETWEEN 3 AND 5
// first_name LIKE 'Bra%' AND last_name <> 'Motley';
//
// @usage
// where([]string{}, []string{}) -> where + operator len must same @deprecated
// where([]string, string) @deprecated
// where(string, nil) @deprecated
// where(string)
//
// # OLD CODE
//
//	if _, ok := where.(string); ok && operator == nil {
//		q.Wheres = fmt.Sprintf("WHERE %s", where)
//	}
//
//	if w, ok := where.([]string); ok {
//		var whereClause string
//
//		whereClause += "WHERE "
//
//		if len(w) > 0 {
//			for i, data := range w {
//				if i != 0 {
//					if _, ok := operator.(string); ok {
//						whereClause += fmt.Sprintf("%s ", operator)
//					}
//
//					if _, ok := operator.([]string); ok {
//						//
//					}
//				}
//
//				whereClause += fmt.Sprintf(`%s `, data)
//			}
//		}
//
//		q.Wheres = whereClause
//	}
func (q SQLSelectBuilder) Where(where string) SQLSelectBuilder {
	q.Wheres += combineConditions()

	return q
}

func (q SQLSelectBuilder) WhereNot(where string) SQLSelectBuilder {
	q.Wheres += combineConditions()

	return q
}

func (q SQLSelectBuilder) Or(where string) SQLSelectBuilder {
	q.Wheres += combineConditions()

	return q
}
