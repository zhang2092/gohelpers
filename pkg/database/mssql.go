package database

import (
	mssql "github.com/denisenkom/go-mssqldb"
)

// MSIsUniqueViolation 唯一值是否冲突 ture: 是, false: 否
func MSIsUniqueViolation(err error) bool {
	msErr, ok := err.(mssql.Error)
	return ok && msErr.Number == 2627
}

//func MsIsForeignKeyViolation(err error) bool {
//	msErr, ok := err.(mssql.Error)
//	return ok && msErr.Number == 0
//}
