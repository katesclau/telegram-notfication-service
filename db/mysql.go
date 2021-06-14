package db

import (
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func buildMysqlDSUrl(url string, database string, username string, password string) string {
	builder := strings.Builder{}
	builder.WriteString(username)
	builder.WriteString(":")
	builder.WriteString(password)
	builder.WriteString("@tcp(")
	builder.WriteString(url)
	builder.WriteString(")/")
	builder.WriteString(database)
	builder.WriteString("?charset=utf8&parseTime=True&loc=Local")

	return builder.String()
}

func NewMYSQLDB(url string, database string, username string, password string) (*gorm.DB, error) {
	dsn := buildMysqlDSUrl(url, database, username, password)
	return gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
}
