package datetime

import (
	"fmt"
	"time"
)

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%04d-%02d-%02d"`, d.Year, d.Month, d.Day)), nil
}

func Today() Date {
	now := time.Now()
	year, month, day := now.Date()
	return Date{
		Year:  year,
		Month: month,
		Day:   day,
	}
}
