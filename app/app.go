package app

import (
	"database/sql"
)

const (
	SiteName = "146ml"
	//RootURL  = "http://146ml.ru"
	RootURL = "http://146ml.ru"
)

var (
	DB *sql.DB
)
