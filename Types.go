package idenfy

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Int returns and Idenfy int type that can be null
func Int(value int) *int {
	return &value
}

// Bool returns and Idenfy bool that can be null
func Bool(value bool) *bool {
	return &value
}

// Date represents a date without time
type Date struct {
	Year  int
	Month int
	Day   int
}

var ErrInvalidDate = errors.New("invalid date")

// MarshalJSON converts a Date into an Idenfy date
func (d Date) MarshalJSON() ([]byte, error) {
	date := fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, ErrInvalidDate
	}
	return []byte(`"` + date + `"`), nil
}

// UnmarshalJSON converts an Idenfy date into a Date
func (d *Date) UnmarshalJSON(data []byte) error {
	date := strings.Trim(string(data), "\"")
	parsedTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		return err
	}
	d.Year = parsedTime.Year()
	d.Month = int(parsedTime.Month())
	d.Day = parsedTime.Day()
	return nil
}
