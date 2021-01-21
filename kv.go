package app

import (
	"database/sql"
	"log"
)

func KVString(k string, db *sql.DB) (values []string) {
	// assumes postgresql syntax and specific table
	dbRows, err := db.Query("select string_value from common.kv where key = $1", k)
	if err != nil {
		log.Println("Error in KVString():", err)
		return
	}
	defer dbRows.Close()
	for dbRows.Next() {
		var s string
		dbRows.Scan(&s)
		values = append(values, s)
	}
	return
}

func KVFloat(k string, db *sql.DB) (values []float64) {
	// assumes postgresql syntax and specific table
	dbRows, err := db.Query("select float_value from common.kv where key = $1", k)
	if err != nil {
		log.Println("Error in KVFloat():", err)
		return
	}
	defer dbRows.Close()
	for dbRows.Next() {
		var f float64
		dbRows.Scan(&f)
		values = append(values, f)
	}
	return
}
