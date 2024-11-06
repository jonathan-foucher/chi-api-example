package models

import (
	"fmt"
	"time"
)

type Date time.Time

func (t *Date) UnmarshalJSON(data []byte) (err error) {
    date, err := time.Parse(`"` + time.DateOnly + `"`, string(data))
    *t = Date(date)
    return
}

func (t Date) MarshalJSON() ([]byte, error) {
	b := fmt.Sprintf(`"` + time.Time(t).Format(time.DateOnly) + `"`)
    return []byte(b), nil
}

func (t Date) String() string {
    return time.Time(t).Format(time.DateOnly)
}
