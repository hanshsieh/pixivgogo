package datetime

import (
	"encoding/json"
	"fmt"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(b []byte) error {
	var timeStr string
	err := json.Unmarshal(b, &timeStr)
	if err != nil {
		return err
	}
	newTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return err
	}
	*t = Time{newTime}
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	result := t.Format(time.RFC3339)
	return []byte(fmt.Sprintf(`"%s"`, result)), nil
}
