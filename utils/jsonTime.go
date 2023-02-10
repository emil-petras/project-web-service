package utils

import (
	"fmt"
	"strconv"
	"time"
)

type JSONTime time.Time

func (t *JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(*t).Format(time.RFC3339))), nil
}

func (t *JSONTime) UnmarshalJSON(b []byte) error {
	unixMilliseconds, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return err
	}

	parsed := time.Unix(0, int64(unixMilliseconds*1e6))
	*t = JSONTime(parsed)
	return nil
}
