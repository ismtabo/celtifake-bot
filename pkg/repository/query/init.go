package query

import "github.com/andreyvit/sqlexpr"

func init() {
	sqlexpr.SetDialect(sqlexpr.PostgresDialect)
}
