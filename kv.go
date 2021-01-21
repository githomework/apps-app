package app

import (
	"database/sql"
	"log"
	"strings"
)

func KVString(k string, db *sql.DB) (values []string) {
	// assumes postgresql syntax and specific table
	dbRows, err := db.Query("select string_value from common.kv where key = $1 and string_value is not null", k)
	if err != nil {
		log.Println("Error in KVString():", err)
		return
	}
	defer dbRows.Close()
	for dbRows.Next() {
		var s sql.NullString
		dbRows.Scan(&s)
		s.String = strings.TrimSpace(s.String)
		if s.String != "" {
			values = append(values, s.String)
		}
	}
	return
}

func KVFloat(k string, db *sql.DB) (values []float64) {
	// assumes postgresql syntax and specific table
	dbRows, err := db.Query("select float_value from common.kv where key = $1 and float_value is not null", k)
	if err != nil {
		log.Println("Error in KVFloat():", err)
		return
	}
	defer dbRows.Close()
	for dbRows.Next() {
		var f sql.NullFloat64
		dbRows.Scan(&f)
		if f.Valid {
			values = append(values, f.Float64)
		}
	}
	return
}
