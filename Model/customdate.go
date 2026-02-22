package model

import (
	"fmt"
	"time"
)

type CustomDate time.Time

const customDateLayout = "02-01-2006"

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1]

	customDate, err := time.Parse(customDateLayout, s)
	if err != nil {
		return fmt.Errorf("invalid date format: %v", err)
	}
	*cd = CustomDate(customDate) //type conversion
	return nil
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	t := time.Time(cd)
	return []byte(`"` + t.Format(customDateLayout) + `"`), nil
}

// // Optional: Add ToTime() if you want to use time.Time later
// func (cd CustomDate) ToTime() time.Time {
// 	return time.Time(cd)
// }
