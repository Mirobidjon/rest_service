package helper

import (
	"database/sql"

	"github.com/golang/protobuf/ptypes/wrappers"
)

func ToNullString(s *wrappers.StringValue) (res sql.NullString) {
	if s.GetValue() != "" {
		res.String = s.Value
		res.Valid = true
	}
	return res
}

func ToStringValue(s sql.NullString) *wrappers.StringValue {
	if s.Valid {
		return &wrappers.StringValue{Value: s.String}
	}
	return nil
}

func ToNullInt(s *wrappers.Int32Value) (res sql.NullInt32) {
	if s.GetValue() != 0 {
		res.Int32 = s.Value
		res.Valid = true
	}
	return res
}

func ToNullInt64(s *wrappers.Int64Value) (res sql.NullInt64) {
	if s.GetValue() != 0 {
		res.Int64 = s.Value
		res.Valid = true
	}
	return res
}
func ToIntValue(s sql.NullInt32) *wrappers.Int32Value {
	if s.Valid {
		return &wrappers.Int32Value{Value: s.Int32}
	}
	return nil
}

func ToInt64Value(s sql.NullInt64) *wrappers.Int64Value {
	if s.Valid {
		return &wrappers.Int64Value{Value: s.Int64}
	}
	return nil
}
