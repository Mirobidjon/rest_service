package config

import "time"

const (
	DatabaseQueryTimeLayout = `'YYYY-MM-DD"T"HH24:MI:SS"."MS"Z"TZ'`
	// DatabaseTimeLayout
	DatabaseTimeLayout string = time.RFC3339
)
