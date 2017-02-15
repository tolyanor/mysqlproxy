package main

import(
	"strconv"
  "database/sql"
  "strings"
)

func exec(sqlString string, db *sql.DB) (string, error) {
	stmt, err := db.Prepare(sqlString)
  if err != nil { return "", err }

  res, err := stmt.Exec()
  if err != nil { return "", err }

  var intRes int64 = -1
	if strings.Index(sqlString, "insert") == 0 {
		intRes, _ = res.LastInsertId()
	} else {
		intRes, _ = res.RowsAffected()
	} 

	return strconv.Itoa(int(intRes)), nil
}