package models

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

func getNullBool(val bool) *sql.NullBool {
	return &sql.NullBool{
		Bool:  val,
		Valid: true,
	}
}

func getNullInt(val int) *sql.NullInt32 {
	if val == 0 {
		return &sql.NullInt32{
			Valid: false,
		}
	} else {
		return &sql.NullInt32{
			Int32: int32(val),
			Valid: true,
		}
	}
}

func getNullString(val string) *sql.NullString {
	if val == "" {
		return &sql.NullString{
			Valid: false,
		}
	} else {
		return &sql.NullString{
			String: val,
			Valid:  true,
		}
	}
}

func getNullTime(val *time.Time) *mysql.NullTime {
	if val == nil || val.IsZero() {
		return &mysql.NullTime{
			Valid: false,
		}
	} else {
		return &mysql.NullTime{
			Time:  *val,
			Valid: true,
		}
	}
}
