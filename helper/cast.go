package helper

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

// NullInt64 type
type NullInt64 sql.NullInt64

// Scan implements the Scanner interface for NullInt64
func (ni *NullInt64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullInt64{i.Int64, false}
	} else {
		*ni = NullInt64{i.Int64, true}
	}
	return nil
}

// MarshalJSON for NullInt64
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON for NullInt64
func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int64)
	ni.Valid = (err == nil)
	return err
}

// ToNullInt64 convert int64 to sql.NullInt64
func ToNullInt64(i int64) sql.NullInt64 {
	//i, err := strconv.Atoi(s)
	return sql.NullInt64{Int64: i, Valid: true}
}

// ToNullString convert
func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}
