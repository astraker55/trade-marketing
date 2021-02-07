package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Layout is using for formating dates
const Layout = "2006-01-02"

// CustomDate is using for 'YYYY-MM-DD' format
type CustomDate struct {
	time.Time
}

// UnmarshalJSON is a method of CustomData to unmarshal this type
func (c *CustomDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	if s == "null" {
		return
	}
	c.Time, err = time.Parse(Layout, s)
	if err != nil {
		return
	}
	return
}

// MarshalJSON is a method of CustomData to unmarshal this type
func (c CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, c.Time.Format(Layout))), nil
}

// Value is using for format database result in CustomDate
func (c *CustomDate) Value() (driver.Value, error) {
	// MyTime is converted to time.Time type
	tTime := time.Time(c.Time)
	return tTime.Format(Layout), nil
}

//Scan is implementing scan interface for CustomDate field
func (c *CustomDate) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		*c = CustomDate{vt}
	default:
		return errors.New("Type handling error")
	}
	return nil
}
