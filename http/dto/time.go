package dto

import (
	"encoding/json"
	"time"
)

const timeFormat = "02.01.2006 в 15:04:05"

type JoinTime struct {
	time.Time
}

func (jt JoinTime) MarshalJSON() ([]byte, error) {
	formatted := jt.Format(timeFormat)
	return json.Marshal(formatted)
}
